package game

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/yokaiio/yokai_server/game/db"
	"github.com/yokaiio/yokai_server/internal/utils"
)

type Game struct {
	sync.RWMutex
	ctx       context.Context
	cancel    context.CancelFunc
	opts      *Options
	waitGroup utils.WaitGroupWrapper

	db      *db.Datastore
	httpSrv *HttpServer
}

func New(opts *Options) (*Game, error) {
	g := &Game{
		opts: opts,
	}

	g.ctx, g.cancel = context.WithCancel(context.Background())
	g.db = db.NewDatastore(g.ctx, opts.MysqlDSN, opts.GameID)
	g.httpSrv = NewHttpServer(g)

	return g, nil
}

func (g *Game) Main() error {

	exitCh := make(chan error)
	var once sync.Once
	exitFunc := func(err error) {
		once.Do(func() {
			if err != nil {
				log.Fatal("Game Main() error:", err)
			}
			exitCh <- err
		})
	}

	// game run
	g.waitGroup.Wrap(func() {
		exitFunc(g.Run())
	})

	// database run
	g.waitGroup.Wrap(func() {
		exitFunc(g.db.Run())
	})

	// http server run
	g.waitGroup.Wrap(func() {
		exitFunc(g.httpSrv.Run())
	})

	err := <-exitCh
	return err
}

func (g *Game) Exit() {
	g.cancel()
	g.waitGroup.Wait()
}

func (g *Game) Run() error {

	for {
		select {
		case <-l.ctx.Done():
			return nil
		default:
		}

		// todo game logic

		t := time.Now()
		d := time.Since(t)
		time.Sleep(time.Second - d)
	}
}
