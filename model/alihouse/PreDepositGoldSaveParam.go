package alihouse

import (
	"sync"
)

// PreDepositGoldSaveParam 结构体
type PreDepositGoldSaveParam struct {
	// 商品标题
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 商品主图
	Pic string `json:"pic,omitempty" xml:"pic,omitempty"`
	// 有效期起-绝对时间 (yyyy-MM-dd)
	ValidStartDate string `json:"valid_start_date,omitempty" xml:"valid_start_date,omitempty"`
	// 有效期止-绝对时间 (yyyy-MM-dd)
	ValidEndDate string `json:"valid_end_date,omitempty" xml:"valid_end_date,omitempty"`
	// 上架时间（yyyy-MM-dd）
	Starts string `json:"starts,omitempty" xml:"starts,omitempty"`
	// 外部项目店id
	OuterStoreId string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
	// 预存金淘宝商品Id,修改时必填
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 有效期-相对时间
	ValidDays int64 `json:"valid_days,omitempty" xml:"valid_days,omitempty"`
	// 购物金价格（单位:分）
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
	// 库存
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 商品状态（立即上架传0 定时上架传1 下架放入仓库传-2
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 是否是拍下减库存，默认是支付减库存
	SubStockAtBuy bool `json:"sub_stock_at_buy,omitempty" xml:"sub_stock_at_buy,omitempty"`
}

var poolPreDepositGoldSaveParam = sync.Pool{
	New: func() any {
		return new(PreDepositGoldSaveParam)
	},
}

// GetPreDepositGoldSaveParam() 从对象池中获取PreDepositGoldSaveParam
func GetPreDepositGoldSaveParam() *PreDepositGoldSaveParam {
	return poolPreDepositGoldSaveParam.Get().(*PreDepositGoldSaveParam)
}

// ReleasePreDepositGoldSaveParam 释放PreDepositGoldSaveParam
func ReleasePreDepositGoldSaveParam(v *PreDepositGoldSaveParam) {
	v.Name = ""
	v.Pic = ""
	v.ValidStartDate = ""
	v.ValidEndDate = ""
	v.Starts = ""
	v.OuterStoreId = ""
	v.ItemId = 0
	v.ValidDays = 0
	v.Price = 0
	v.Quantity = 0
	v.Status = 0
	v.SubStockAtBuy = false
	poolPreDepositGoldSaveParam.Put(v)
}
