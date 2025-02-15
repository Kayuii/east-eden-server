package att

import (
	"fmt"
	"testing"

	"github.com/east-eden/server/define"
	"github.com/east-eden/server/excel"
	"github.com/east-eden/server/utils"
	"github.com/shopspring/decimal"
)

func TestAttManager(t *testing.T) {
	// reload to project root path
	if err := utils.RelocatePath("/server"); err != nil {
		t.Fatalf("TestAttManager failed: %s", err.Error())
	}

	excel.ReadAllEntries("config/csv/")

	attManager := NewAttManager()
	attManager.SetBaseAttId(1)
	// attManager.ModBaseAtt(define.Att_Atk, 100)

	attManager2 := NewAttManager()
	attManager2.SetBaseAttId(2)
	attManager.ModAttManager(attManager2)
	attManager.CalcAtt()
	_ = attManager.GetFinalAttValue(define.Att_AtkBase)

	d1, _ := decimal.NewFromString("101.57")
	d2, _ := decimal.NewFromString("-382.4")
	x := d1.Round(0).IntPart()
	fmt.Println(x)
	y := d2.IntPart()
	fmt.Println(y)
	round := int32(d1.Round(0).IntPart())
	fmt.Println(round)
	d3 := d1.Mul(d2)
	r := d3.Floor().BigInt().Int64()
	fmt.Println(r)
}
