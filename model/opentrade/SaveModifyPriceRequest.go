package opentrade

import (
	"sync"
)

// SaveModifyPriceRequest 结构体
type SaveModifyPriceRequest struct {
	// 买家openId，如果有就传，后续会校验。没有的可以不用传，但是校验买家参数会跳过
	OpenId string `json:"open_id,omitempty" xml:"open_id,omitempty"`
	// 改价的商品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 改价价格，单位分
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
}

var poolSaveModifyPriceRequest = sync.Pool{
	New: func() any {
		return new(SaveModifyPriceRequest)
	},
}

// GetSaveModifyPriceRequest() 从对象池中获取SaveModifyPriceRequest
func GetSaveModifyPriceRequest() *SaveModifyPriceRequest {
	return poolSaveModifyPriceRequest.Get().(*SaveModifyPriceRequest)
}

// ReleaseSaveModifyPriceRequest 释放SaveModifyPriceRequest
func ReleaseSaveModifyPriceRequest(v *SaveModifyPriceRequest) {
	v.OpenId = ""
	v.ItemId = 0
	v.Price = 0
	poolSaveModifyPriceRequest.Put(v)
}
