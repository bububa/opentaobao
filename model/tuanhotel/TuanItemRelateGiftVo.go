package tuanhotel

import (
	"sync"
)

// TuanItemRelateGiftVo 结构体
type TuanItemRelateGiftVo struct {
	// 单位
	Unit string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 描述
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 数量
	Num int64 `json:"num,omitempty" xml:"num,omitempty"`
	// 类别：1-温泉；2-滑雪；3-门票；4-美食；5-下午茶；6-SPA；7-礼包；8-亲子；9-交通；10-代金券；999-用户自定义
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 宝贝ID
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
}

var poolTuanItemRelateGiftVo = sync.Pool{
	New: func() any {
		return new(TuanItemRelateGiftVo)
	},
}

// GetTuanItemRelateGiftVo() 从对象池中获取TuanItemRelateGiftVo
func GetTuanItemRelateGiftVo() *TuanItemRelateGiftVo {
	return poolTuanItemRelateGiftVo.Get().(*TuanItemRelateGiftVo)
}

// ReleaseTuanItemRelateGiftVo 释放TuanItemRelateGiftVo
func ReleaseTuanItemRelateGiftVo(v *TuanItemRelateGiftVo) {
	v.Unit = ""
	v.Desc = ""
	v.Num = 0
	v.Type = 0
	v.ItemId = 0
	poolTuanItemRelateGiftVo.Put(v)
}
