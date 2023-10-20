package travel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelitemsinglequeryAPIRequest 【API3.0】度假单个商品查询接口 API请求
// taobao.alitrip.travel.item.single.query
//
// 旅行度假新商品查询接口（单个商品查询） 第三版
type TaobaoalitriptravelitemsinglequeryAPIRequest struct {
	model.Params
	// 商品 外部商家编码。itemId和outProductId至少填写一个
	_outProductId string
	// 商品id。itemId和outProductId至少填写一个
	_itemId int64
}

// NewTaobaoalitriptravelitemsinglequeryRequest 初始化TaobaoalitriptravelitemsinglequeryAPIRequest对象
func NewTaobaoalitriptravelitemsinglequeryRequest() *TaobaoalitriptravelitemsinglequeryAPIRequest {
	return &TaobaoalitriptravelitemsinglequeryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitriptravelitemsinglequeryAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.item.single.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitriptravelitemsinglequeryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitriptravelitemsinglequeryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOutProductId is OutProductId Setter
// 商品 外部商家编码。itemId和outProductId至少填写一个
func (r *TaobaoalitriptravelitemsinglequeryAPIRequest) SetOutProductId(_outProductId string) error {
	r._outProductId = _outProductId
	r.Set("out_product_id", _outProductId)
	return nil
}

// GetOutProductId OutProductId Getter
func (r TaobaoalitriptravelitemsinglequeryAPIRequest) GetOutProductId() string {
	return r._outProductId
}

// SetItemId is ItemId Setter
// 商品id。itemId和outProductId至少填写一个
func (r *TaobaoalitriptravelitemsinglequeryAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoalitriptravelitemsinglequeryAPIRequest) GetItemId() int64 {
	return r._itemId
}
