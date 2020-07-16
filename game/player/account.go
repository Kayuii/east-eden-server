package player

import (
	"context"
	"fmt"
	"time"

	"github.com/golang/protobuf/proto"
	logger "github.com/sirupsen/logrus"
	"github.com/yokaiio/yokai_server/define"
	pbAccount "github.com/yokaiio/yokai_server/proto/account"
	"github.com/yokaiio/yokai_server/store"
	"github.com/yokaiio/yokai_server/transport"
	"github.com/yokaiio/yokai_server/utils"
)

var (
	Account_MemExpire = time.Hour * 2
)

// lite account info
type LiteAccount struct {
	store.StoreObjector `bson:"-" json:"-"`
	ID                  int64   `bson:"_id" json:"_id"`
	UserId              int64   `bson:"user_id" json:"user_id"`
	GameId              int16   `bson:"game_id" json:"game_id"`
	Name                string  `bson:"name" json:"name"`
	Level               int32   `bson:"level" json:"level"`
	PlayerIDs           []int64 `bson:"player_id" json:"player_id"`
}

func (la *LiteAccount) GetObjID() int64 {
	return la.ID
}

func (la *LiteAccount) GetStoreIndex() int64 {
	return -1
}

func (la *LiteAccount) AfterLoad() error {
	return nil
}

func (la *LiteAccount) GetID() int64 {
	return la.ID
}

func (la *LiteAccount) SetID(id int64) {
	la.ID = id
}

func (la *LiteAccount) GetName() string {
	return la.Name
}

func (la *LiteAccount) SetName(name string) {
	la.Name = name
}

func (la *LiteAccount) GetLevel() int32 {
	return la.Level
}

func (la *LiteAccount) SetLevel(level int32) {
	la.Level = level
}

func (la *LiteAccount) AddPlayerID(playerID int64) {
	for _, value := range la.PlayerIDs {
		if value == playerID {
			return
		}
	}

	la.PlayerIDs = append(la.PlayerIDs, playerID)
}

func (la *LiteAccount) GetPlayerIDs() []int64 {
	return la.PlayerIDs
}

// full account info
type Account struct {
	LiteAccount `bson:"inline" json:",inline"`

	sock transport.Socket `bson:"-" json:"-"`
	p    *Player          `bson:"-" json:"-"`

	waitGroup utils.WaitGroupWrapper `bson:"-" json:"-"`
	timeOut   *time.Timer            `bson:"-" json:"-"`

	chHandler chan func(*Account) `bson:"-" json:"-"`
}

func NewLiteAccount() interface{} {
	return &LiteAccount{
		ID:        -1,
		Name:      "",
		Level:     1,
		PlayerIDs: []int64{},
	}
}

func NewAccount() interface{} {
	account := &Account{
		LiteAccount: *(NewLiteAccount().(*LiteAccount)),
		sock:        nil,
		p:           nil,
		timeOut:     time.NewTimer(define.Account_OnlineTimeout),
	}

	return account
}

func (a *Account) GetSock() transport.Socket {
	return a.sock
}

func (a *Account) SetSock(s transport.Socket) {
	a.sock = s
}

func (a *Account) SetLaterHandlerChannel(ch chan func(*Account)) {
	a.chHandler = ch
}

func (a *Account) GetPlayer() *Player {
	return a.p
}

func (a *Account) SetPlayer(p *Player) {
	a.p = p
}

func (a *Account) Main(ctx context.Context) error {

	a.waitGroup.Wrap(func() {
		a.Run(ctx)
	})

	a.waitGroup.Wait()

	return nil
}

func (a *Account) Exit() {
	a.timeOut.Stop()
	a.sock.Close()
}

func (a *Account) Run(ctx context.Context) error {
	for {
		select {
		// context canceled
		case <-ctx.Done():
			logger.WithFields(logger.Fields{
				"id":            a.GetID(),
				"socket_remote": a.sock.Remote(),
			}).Info("Account context done!")
			return nil

		case h, ok := <-a.chHandler:
			if !ok {
				continue
			}

			h(a)

		// lost connection
		case <-a.timeOut.C:
			return fmt.Errorf("account<%d> time out", a.GetID())
		}
	}
}

/*
msg Example:
	Type: transport.BodyProtobuf
	Name: yokai_account.M2C_AccountLogon
	Body: protoBuf byte
*/
func (a *Account) SendProtoMessage(p proto.Message) {
	if a.sock == nil {
		return
	}

	var msg transport.Message
	msg.Type = transport.BodyProtobuf
	msg.Name = proto.MessageName(p)
	msg.Body = p

	if err := a.sock.Send(&msg); err != nil {
		logger.Warn("send proto msg error:", err)
		return
	}
}

func (a *Account) HeartBeat(rpcId int32) {
	a.timeOut.Reset(define.Account_OnlineTimeout)

	reply := &pbAccount.M2C_HeartBeat{RpcId: rpcId, Timestamp: uint32(time.Now().Unix())}
	a.SendProtoMessage(reply)
}
