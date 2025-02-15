package rank

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/asim/go-micro/v3/client"
	"github.com/east-eden/server/define"
	pbGlobal "github.com/east-eden/server/proto/global"
	pbGame "github.com/east-eden/server/proto/server/game"
	pbRank "github.com/east-eden/server/proto/server/rank"
	"github.com/east-eden/server/utils"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

var (
	ErrInvalidGlobalConfig = errors.New("invalid global config")
)

var (
	DefaultRpcTimeout = 5 * time.Second // 默认rpc超时时间
)

type RpcHandler struct {
	m       *Rank
	rankSrv pbRank.RankService
	gameSrv pbGame.GameService
}

func NewRpcHandler(cli *cli.Context, m *Rank) *RpcHandler {
	h := &RpcHandler{
		m: m,
		rankSrv: pbRank.NewRankService(
			"rank",
			m.mi.srv.Client(),
		),
		gameSrv: pbGame.NewGameService(
			"game",
			m.mi.srv.Client(),
		),
	}

	err := pbRank.RegisterRankServiceHandler(m.mi.srv.Server(), h)
	if err != nil {
		log.Fatal().Err(err).Msg("RegisterRankServiceHandler failed")
	}

	return h
}

// 一致性哈希
func (h *RpcHandler) consistentHashCallOption(key string) client.CallOption {
	return client.WithSelectOption(
		utils.ConsistentHashSelector(h.m.cons, key),
	)
}

// 重试次数
func (h *RpcHandler) retries(times int) client.CallOption {
	return client.WithRetries(times)
}

/////////////////////////////////////////////
// rpc call
/////////////////////////////////////////////
func (h *RpcHandler) CallKickRankData(rankId int32, nodeId int32) (*pbRank.KickRankDataRs, error) {
	if rankId == -1 {
		return nil, ErrInvalidRank
	}

	if nodeId == int32(h.m.ID) {
		return nil, errors.New("same rank node id")
	}

	req := &pbRank.KickRankDataRq{
		RankId:     rankId,
		RankNodeId: nodeId,
	}

	ctx, cancel := context.WithTimeout(context.Background(), DefaultRpcTimeout)
	defer cancel()

	return h.rankSrv.KickRankData(
		ctx,
		req,
		client.WithSelectOption(
			utils.SpecificIDSelector(
				fmt.Sprintf("rank-%d", nodeId),
			),
		),
	)
}

/////////////////////////////////////////////
// rpc receive
/////////////////////////////////////////////

// 查询排行
func (h *RpcHandler) QueryRankByObjId(
	ctx context.Context,
	req *pbRank.QueryRankByObjIdRq,
	rsp *pbRank.QueryRankByObjIdRs,
) error {
	rsp.RankId = req.GetRankId()
	rsp.ObjId = req.GetObjId()

	rank, metadata, err := h.m.manager.QueryRankByObjId(ctx, req.GetRankId(), req.GetObjId())
	if utils.ErrCheck(err, "QueryRankByObjId failed when RpcHandler.QueryRankByObjId") {
		rsp.RankIndex = int32(rank)
		rsp.Metadata = metadata.ToPB()
	}
	return err
}

func (h *RpcHandler) QueryRankByRange(
	ctx context.Context,
	req *pbRank.QueryRankByRangeRq,
	rsp *pbRank.QueryRankByRangeRs,
) error {
	rsp.RankId = req.GetRankId()
	metadatas, err := h.m.manager.QueryRankByRange(ctx, req.GetRankId(), req.GetStart(), req.GetEnd())
	rsp.Metadatas = make([]*pbGlobal.RankMetadata, 0, len(metadatas))
	for _, metadata := range metadatas {
		rsp.Metadatas = append(rsp.Metadatas, metadata.ToPB())
	}
	return err
}

// 踢出邮件cache
func (h *RpcHandler) KickRankData(
	ctx context.Context,
	req *pbRank.KickRankDataRq,
	rsp *pbRank.KickRankDataRs,
) error {
	return h.m.manager.KickRankData(req.GetRankId(), req.GetRankNodeId())
}

// 设置排行积分
func (h *RpcHandler) SetRankScore(
	ctx context.Context,
	req *pbRank.SetRankScoreRq,
	rsp *pbRank.SetRankScoreRs,
) error {
	metadata := &define.RankMetadata{
		RankKey: define.RankKey{
			ObjId:  req.GetMetadata().ObjId,
			RankId: req.GetRankId(),
		},
	}
	metadata.FromPB(req.GetMetadata())
	return h.m.manager.SetRankScore(ctx, req.GetRankId(), metadata)
}
