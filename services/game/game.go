package game

import (
	"context"
	"fmt"
	"os"
	"sync"

	"bitbucket.org/funplus/server/excel"
	"bitbucket.org/funplus/server/logger"
	pbGlobal "bitbucket.org/funplus/server/proto/global"
	"bitbucket.org/funplus/server/services/game/global"
	"bitbucket.org/funplus/server/store"
	"bitbucket.org/funplus/server/utils"
	"github.com/rs/zerolog"
	log "github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
	"stathat.com/c/consistent"
)

var (
	maxGameNode = 128 // max game node number, used in constent hash
)

type Game struct {
	app *cli.App
	ID  int16
	sync.RWMutex
	wg utils.WaitGroupWrapper

	tcpSrv      *TcpServer
	wsSrv       *WsServer
	gin         *GinServer
	am          *AccountManager
	mi          *MicroService
	rpcHandler  *RpcHandler
	msgRegister *MsgRegister
	pubSub      *PubSub
	cons        *consistent.Consistent
}

func New() *Game {
	g := &Game{}

	g.app = cli.NewApp()
	g.app.Name = "game"
	g.app.Flags = NewFlags()

	g.app.Before = g.Before
	g.app.Action = g.Action
	g.app.UsageText = "game [first_arg] [second_arg]"
	g.app.Authors = []*cli.Author{{Name: "dudu", Email: "hellodudu86@gmail"}}

	return g
}

func (g *Game) Before(ctx *cli.Context) error {
	// relocate path
	if err := utils.RelocatePath("/server_bin", "\\server_bin", "/server", "\\server"); err != nil {
		fmt.Println("relocate failed: ", err)
		os.Exit(1)
	}

	// logger init
	logger.InitLogger("game")

	// load excel entries
	excel.ReadAllEntries("config/excel/")

	// read config/game/config.toml
	return altsrc.InitInputSourceWithContext(g.app.Flags, altsrc.NewTomlSourceFromFlagFunc("config_file"))(ctx)
}

func (g *Game) Action(ctx *cli.Context) error {
	// logger settings
	logLevel, err := zerolog.ParseLevel(ctx.String("log_level"))
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	log.Logger = log.Level(logLevel)

	exitCh := make(chan error)
	var once sync.Once
	exitFunc := func(err error) {
		once.Do(func() {
			if err != nil {
				log.Fatal().Err(err).Msg("Game Action() failed")
			}
			exitCh <- err
		})
	}

	g.ID = int16(ctx.Int("game_id"))

	// init snowflakes
	utils.InitMachineID(g.ID)

	store.NewStore(ctx)
	g.am = NewAccountManager(ctx, g)
	g.gin = NewGinServer(ctx, g)
	g.mi = NewMicroService(ctx, g)
	g.rpcHandler = NewRpcHandler(g)
	g.pubSub = NewPubSub(g)
	g.msgRegister = NewMsgRegister(g.am, g.rpcHandler, g.pubSub)
	g.tcpSrv = NewTcpServer(ctx, g)
	g.wsSrv = NewWsServer(ctx, g)
	g.cons = consistent.New()
	g.cons.NumberOfReplicas = maxGameNode

	// tcp server run
	g.wg.Wrap(func() {
		defer utils.CaptureException()
		exitFunc(g.tcpSrv.Run(ctx))
		g.tcpSrv.Exit()
	})

	// websocket server
	g.wg.Wrap(func() {
		defer utils.CaptureException()
		exitFunc(g.wsSrv.Run(ctx))
		g.wsSrv.Exit()
	})

	// gin server
	g.wg.Wrap(func() {
		defer utils.CaptureException()
		exitFunc(g.gin.Main(ctx))
		g.gin.Exit(ctx)
	})

	// client mgr run
	g.wg.Wrap(func() {
		defer utils.CaptureException()
		exitFunc(g.am.Main(ctx))
		g.am.Exit()

	})

	// micro run
	g.wg.Wrap(func() {
		defer utils.CaptureException()
		exitFunc(g.mi.Run())
	})

	// global mess run
	g.wg.Wrap(func() {
		defer utils.CaptureException()
		exitFunc(global.GetGlobalMess().Run(ctx))
	})

	return <-exitCh
}

func (g *Game) Run(arguments []string) error {

	// app run
	if err := g.app.Run(arguments); err != nil {
		return err
	}

	return nil
}

func (g *Game) Stop() {
	store.GetStore().Exit()
	g.wg.Wait()
}

///////////////////////////////////////////////////////
// pubsub
///////////////////////////////////////////////////////
func (g *Game) StartGate() {
	srvs, _ := g.mi.srv.Server().Options().Registry.GetService("game")
	for _, v := range srvs {
		log.Info().Str("name", v.Name).Msg("list all services")
		for _, n := range v.Nodes {
			log.Info().Interface("node", n).Msg("list nodes")
		}
	}

	c := &pbGlobal.AccountInfo{Id: 12, Name: "game's client 12"}
	err := g.pubSub.PubStartGate(context.Background(), c)
	log.Info().Err(err).Msg("publish start gate result")
}
