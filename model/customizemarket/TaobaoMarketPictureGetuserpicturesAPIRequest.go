package customizemarket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoMarketPictureGetuserpicturesAPIRequest
读取用户上传图片 API请求
taobao.market.picture.getuserpictures

商家通过用户信息，获取用户上传的 */
type TaobaoMarketPictureGetuserpicturesAPIRequest struct {
	model.Params
	// 订单ID
	_orderId int64
}

// NewTaobaoMarketPictureGetuserpicturesRequest 初始化TaobaoMarketPictureGetuserpicturesAPIRequest对象
func NewTaobaoMarketPictureGetuserpicturesRequest() *TaobaoMarketPictureGetuserpicturesAPIRequest {
	return &TaobaoMarketPictureGetuserpicturesAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMarketPictureGetuserpicturesAPIRequest) GetApiMethodName() string {
	return "taobao.market.picture.getuserpictures"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMarketPictureGetuserpicturesAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OrderId Setter
// 订单ID
func (r *TaobaoMarketPictureGetuserpicturesAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// Get OrderId Getter
func (r TaobaoMarketPictureGetuserpicturesAPIRequest) GetOrderId() int64 {
	return r._orderId
}
