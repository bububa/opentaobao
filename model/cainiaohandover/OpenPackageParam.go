package cainiaohandover

import (
	"sync"
)

// OpenPackageParam 结构体
type OpenPackageParam struct {
	// 商品参数
	ItemParams []OpenItemParam `json:"item_params,omitempty" xml:"item_params>open_item_param,omitempty"`
	// 包裹价格币种，CNY：人民币、USD：美元、RUB：卢布。
	Currency string `json:"currency,omitempty" xml:"currency,omitempty"`
	// 包裹长度
	Length int64 `json:"length,omitempty" xml:"length,omitempty"`
	// 包裹宽度
	Width int64 `json:"width,omitempty" xml:"width,omitempty"`
	// 包裹高度
	Height int64 `json:"height,omitempty" xml:"height,omitempty"`
	// 包裹重量
	Weight int64 `json:"weight,omitempty" xml:"weight,omitempty"`
	// 包裹价格
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
}

var poolOpenPackageParam = sync.Pool{
	New: func() any {
		return new(OpenPackageParam)
	},
}

// GetOpenPackageParam() 从对象池中获取OpenPackageParam
func GetOpenPackageParam() *OpenPackageParam {
	return poolOpenPackageParam.Get().(*OpenPackageParam)
}

// ReleaseOpenPackageParam 释放OpenPackageParam
func ReleaseOpenPackageParam(v *OpenPackageParam) {
	v.ItemParams = v.ItemParams[:0]
	v.Currency = ""
	v.Length = 0
	v.Width = 0
	v.Height = 0
	v.Weight = 0
	v.Price = 0
	poolOpenPackageParam.Put(v)
}
