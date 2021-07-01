package travel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripTravelItemNewQueryAPIRequest
【API3.0】新版度假单个商品查询接口 API请求
taobao.alitrip.travel.item.new.query

新版旅行度假新商品查询接口（单个商品查询） */
type TaobaoAlitripTravelItemNewQueryAPIRequest struct {
	model.Params
	// 商品id。itemId和outProductId至少填写一个
	_itemId int64
	// 商品 外部商家编码。itemId和outProductId至少填写一个
	_outProductId string
}

// NewTaobaoAlitripTravelItemNewQueryRequest 初始化TaobaoAlitripTravelItemNewQueryAPIRequest对象
func NewTaobaoAlitripTravelItemNewQueryRequest() *TaobaoAlitripTravelItemNewQueryAPIRequest {
	return &TaobaoAlitripTravelItemNewQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelItemNewQueryAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.item.new.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelItemNewQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ItemId Setter
// 商品id。itemId和outProductId至少填写一个
func (r *TaobaoAlitripTravelItemNewQueryAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// Get ItemId Getter
func (r TaobaoAlitripTravelItemNewQueryAPIRequest) GetItemId() int64 {
	return r._itemId
}

// Set is OutProductId Setter
// 商品 外部商家编码。itemId和outProductId至少填写一个
func (r *TaobaoAlitripTravelItemNewQueryAPIRequest) SetOutProductId(_outProductId string) error {
	r._outProductId = _outProductId
	r.Set("out_product_id", _outProductId)
	return nil
}

// Get OutProductId Getter
func (r TaobaoAlitripTravelItemNewQueryAPIRequest) GetOutProductId() string {
	return r._outProductId
}
