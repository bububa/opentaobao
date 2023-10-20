package customizemarket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaomarketpicturegetuserpicturesAPIRequest 读取用户上传图片 API请求
// taobao.market.picture.getuserpictures
//
// 商家通过用户信息，获取用户上传的
type TaobaomarketpicturegetuserpicturesAPIRequest struct {
	model.Params
	// 订单ID
	_orderId int64
}

// NewTaobaomarketpicturegetuserpicturesRequest 初始化TaobaomarketpicturegetuserpicturesAPIRequest对象
func NewTaobaomarketpicturegetuserpicturesRequest() *TaobaomarketpicturegetuserpicturesAPIRequest {
	return &TaobaomarketpicturegetuserpicturesAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaomarketpicturegetuserpicturesAPIRequest) GetApiMethodName() string {
	return "taobao.market.picture.getuserpictures"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaomarketpicturegetuserpicturesAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaomarketpicturegetuserpicturesAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 订单ID
func (r *TaobaomarketpicturegetuserpicturesAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TaobaomarketpicturegetuserpicturesAPIRequest) GetOrderId() int64 {
	return r._orderId
}
