package game

import (
	logger "github.com/sirupsen/logrus"
	"github.com/yokaiio/yokai_server/internal/transport"
	pbGame "github.com/yokaiio/yokai_server/proto/game"
)

func (m *MsgHandler) handleAddItem(sock transport.Socket, p *transport.Message) {
	acct := m.g.am.GetAccountBySock(sock)
	if acct == nil {
		logger.WithFields(logger.Fields{
			"account_id":   acct.GetID(),
			"account_name": acct.GetName(),
		}).Warn("add item failed")
		return
	}

	pl := m.g.pm.GetPlayerByAccount(acct)
	if pl == nil {
		return
	}

	msg, ok := p.Body.(*pbGame.C2M_AddItem)
	if !ok {
		logger.Warn("Add Item failed, recv message body error")
		return
	}

	acct.PushWrapHandler(func() {
		pl.ItemManager().AddItemByTypeID(msg.TypeId, 1)
		list := pl.ItemManager().GetItemList()
		reply := &pbGame.M2C_ItemList{}
		for _, v := range list {
			i := &pbGame.Item{
				Id:     v.GetID(),
				TypeId: v.GetTypeID(),
			}
			reply.Items = append(reply.Items, i)
		}
		acct.SendProtoMessage(reply)
	})

}

func (m *MsgHandler) handleDelItem(sock transport.Socket, p *transport.Message) {
	acct := m.g.am.GetAccountBySock(sock)
	if acct == nil {
		logger.WithFields(logger.Fields{
			"account_id":   acct.GetID(),
			"account_name": acct.GetName(),
		}).Warn("delete item failed")
		return
	}

	pl := m.g.pm.GetPlayerByAccount(acct)
	if pl == nil {
		return
	}

	msg, ok := p.Body.(*pbGame.C2M_DelItem)
	if !ok {
		logger.Warn("Delete item failed, recv message body error")
		return
	}

	acct.PushWrapHandler(func() {
		item := pl.ItemManager().GetItem(msg.Id)
		if item == nil {
			logger.Warn("Delete item failed, non-existing item_id:", msg.Id)
			return
		}

		// clear hero's equip id before delete item
		equipObjID := item.GetEquipObj()
		if equipObjID != -1 {
			pl.HeroManager().TakeoffEquip(equipObjID, item.EquipEnchantEntry().EquipPos)
		}

		// delete item
		pl.ItemManager().DeleteItem(msg.Id)

		// reply to client
		list := pl.ItemManager().GetItemList()
		reply := &pbGame.M2C_ItemList{}
		for _, v := range list {
			i := &pbGame.Item{
				Id:     v.GetID(),
				TypeId: v.GetTypeID(),
			}
			reply.Items = append(reply.Items, i)
		}
		acct.SendProtoMessage(reply)
	})
}

func (m *MsgHandler) handleUseItem(sock transport.Socket, p *transport.Message) {
	acct := m.g.am.GetAccountBySock(sock)
	if acct == nil {
		logger.WithFields(logger.Fields{
			"account_id":   acct.GetID(),
			"account_name": acct.GetName(),
		}).Warn("use item failed")
		return
	}

	pl := m.g.pm.GetPlayerByAccount(acct)
	if pl == nil {
		return
	}

	msg, ok := p.Body.(*pbGame.C2M_UseItem)
	if !ok {
		logger.Warn("Use Item failed, recv message body error")
		return
	}

	acct.PushWrapHandler(func() {
		pl.ItemManager().UseItem(msg.ItemId)
	})
}

func (m *MsgHandler) handleQueryItems(sock transport.Socket, p *transport.Message) {
	acct := m.g.am.GetAccountBySock(sock)
	if acct == nil {
		logger.WithFields(logger.Fields{
			"account_id":   acct.GetID(),
			"account_name": acct.GetName(),
		}).Warn("query items failed")
		return
	}

	pl := m.g.pm.GetPlayerByAccount(acct)
	if pl == nil {
		return
	}

	acct.PushWrapHandler(func() {
		reply := &pbGame.M2C_ItemList{}
		list := pl.ItemManager().GetItemList()
		for _, v := range list {
			i := &pbGame.Item{
				Id:     v.GetID(),
				TypeId: v.GetTypeID(),
			}
			reply.Items = append(reply.Items, i)
		}
		acct.SendProtoMessage(reply)
	})
}

func (m *MsgHandler) handlePutonEquip(sock transport.Socket, p *transport.Message) {
	acct := m.g.am.GetAccountBySock(sock)
	if acct == nil {
		logger.WithFields(logger.Fields{
			"account_id":   acct.GetID(),
			"account_name": acct.GetName(),
		}).Warn("Puton equip failed")
		return
	}

	pl := m.g.pm.GetPlayerByAccount(acct)
	if pl == nil {
		return
	}

	msg, ok := p.Body.(*pbGame.C2M_PutonEquip)
	if !ok {
		logger.Warn("Puton equip failed, recv message body error")
		return
	}

	acct.PushWrapHandler(func() {
		if err := pl.HeroManager().PutonEquip(msg.HeroId, msg.EquipId); err != nil {
			logger.Warn(err)
			return
		}
	})
}

func (m *MsgHandler) handleTakeoffEquip(sock transport.Socket, p *transport.Message) {
	acct := m.g.am.GetAccountBySock(sock)
	if acct == nil {
		logger.WithFields(logger.Fields{
			"account_id":   acct.GetID(),
			"account_name": acct.GetName(),
		}).Warn("Takeoff equip failed")
		return
	}

	pl := m.g.pm.GetPlayerByAccount(acct)
	if pl == nil {
		return
	}

	msg, ok := p.Body.(*pbGame.C2M_TakeoffEquip)
	if !ok {
		logger.Warn("Takeoff equip failed, recv message body error")
		return
	}

	acct.PushWrapHandler(func() {
		if err := pl.HeroManager().TakeoffEquip(msg.HeroId, msg.Pos); err != nil {
			logger.Warn(err)
			return
		}
	})
}

func (m *MsgHandler) handleQueryHeroEquips(sock transport.Socket, p *transport.Message) {
	acct := m.g.am.GetAccountBySock(sock)
	if acct == nil {
		logger.WithFields(logger.Fields{
			"account_id":   acct.GetID(),
			"account_name": acct.GetName(),
		}).Warn("Query hero equips failed")
		return
	}

	pl := m.g.pm.GetPlayerByAccount(acct)
	if pl == nil {
		return
	}

	msg, ok := p.Body.(*pbGame.C2M_QueryHeroEquips)
	if !ok {
		logger.Warn("Query hero equips failed, recv message body error")
		return
	}

	acct.PushWrapHandler(func() {
		hero := pl.HeroManager().GetHero(msg.HeroId)
		if hero == nil {
			logger.Warn("Query hero equips failed, non-existing hero_id:", msg.HeroId)
			return
		}

		reply := &pbGame.M2C_HeroEquips{
			HeroId: msg.HeroId,
			Equips: make([]*pbGame.Item, 0),
		}

		equips := hero.GetEquips()
		for _, v := range equips {
			if v == -1 {
				continue
			}

			it := pl.ItemManager().GetItem(v)
			i := &pbGame.Item{
				Id:     v,
				TypeId: it.GetTypeID(),
			}
			reply.Equips = append(reply.Equips, i)
		}
		acct.SendProtoMessage(reply)
	})
}
