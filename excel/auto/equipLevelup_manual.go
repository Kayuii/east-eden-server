package auto

import (
	"github.com/east-eden/server/define"
	"github.com/east-eden/server/excel"
	"github.com/east-eden/server/utils"
)

var (
	equipLevelTotalExp map[int32][]int32
)

func init() {
	equipLevelTotalExp = make(map[int32][]int32)
	excel.AddEntryManualLoader("EquipLevelup.xlsx", (*EquipLevelupEntries)(nil))
}

// ManualLoader
func (e *EquipLevelupEntries) ManualLoad(*excel.ExcelFileRaw) error {
	rows := GetEquipLevelupRows()
	for level := range rows {
		_ = GetEquipLevelTotalExp(level, 0)
	}

	return nil
}

// 获取装备某一等级含有的所有经验值
func GetEquipLevelTotalExp(level int32, quality int32) int32 {
	if level <= 0 {
		return 0
	}

	if !utils.Between(quality, define.Item_Quality_Begin, define.Item_Quality_End) {
		return 0
	}

	exps, ok := equipLevelTotalExp[level]
	if ok {
		return exps[quality]
	}

	// 第一次找不到数据的话则生成一份
	equipLevelTotalExp[level] = make([]int32, define.Item_Quality_Num)
	var n int32
	for n = 1; n <= level; n++ {
		levelupEntry, _ := GetEquipLevelupEntry(n)
		for idx := range levelupEntry.Exp {
			equipLevelTotalExp[level][idx] += levelupEntry.Exp[idx]
		}
	}

	return equipLevelTotalExp[level][quality]
}
