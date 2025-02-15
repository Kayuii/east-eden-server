package item

import (
	"fmt"

	"github.com/east-eden/server/define"
	"github.com/east-eden/server/utils"
)

type EquipBar struct {
	owner     define.PluginObj
	equipList [define.Equip_Pos_End]*Equip
}

func NewEquipBar(owner define.PluginObj) *EquipBar {
	m := &EquipBar{
		owner: owner,
	}

	return m
}

func (eb *EquipBar) GetEquipByPos(pos int32) *Equip {
	if !utils.Between(pos, define.Equip_Pos_Begin, define.Equip_Pos_End) {
		return nil
	}

	return eb.equipList[pos]
}

func (eb *EquipBar) PutonEquip(e *Equip) error {
	pos := e.EquipEnchantEntry.EquipPos
	if !utils.Between(pos, define.Equip_Pos_Begin, define.Equip_Pos_End) {
		return fmt.Errorf("puton equip error: invalid pos<%d>", pos)
	}

	if eb.GetEquipByPos(pos) != nil {
		return fmt.Errorf("puton equip error: existing equip on this pos<%d>", pos)
	}

	eb.equipList[pos] = e
	e.EquipObj = eb.owner.GetId()
	return nil
}

func (eb *EquipBar) TakeoffEquip(pos int32) error {
	if !utils.Between(pos, define.Equip_Pos_Begin, define.Equip_Pos_End) {
		return fmt.Errorf("takeoff equip error: invalid pos<%d>", pos)
	}

	if e := eb.equipList[pos]; e != nil {
		e.EquipObj = -1
	}

	eb.equipList[pos] = nil
	return nil
}
