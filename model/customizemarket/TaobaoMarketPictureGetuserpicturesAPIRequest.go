package customizemarket

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMarketPictureGetuserpicturesAPIRequest 读取用户上传图片 API请求
// taobao.market.picture.getuserpictures
//
// 商家通过用户信息，获取用户上传的
type TaobaoMarketPictureGetuserpicturesAPIRequest struct {
	model.Params
	// 订单ID
	_orderId int64
}

// NewTaobaoMarketPictureGetuserpicturesRequest 初始化TaobaoMarketPictureGetuserpicturesAPIRequest对象
func NewTaobaoMarketPictureGetuserpicturesRequest() *TaobaoMarketPictureGetuserpicturesAPIRequest {
	return &TaobaoMarketPictureGetuserpicturesAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoMarketPictureGetuserpicturesAPIRequest) Reset() {
	r._orderId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMarketPictureGetuserpicturesAPIRequest) GetApiMethodName() string {
	return "taobao.market.picture.getuserpictures"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMarketPictureGetuserpicturesAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoMarketPictureGetuserpicturesAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 订单ID
func (r *TaobaoMarketPictureGetuserpicturesAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TaobaoMarketPictureGetuserpicturesAPIRequest) GetOrderId() int64 {
	return r._orderId
}

var poolTaobaoMarketPictureGetuserpicturesAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoMarketPictureGetuserpicturesRequest()
	},
}

// GetTaobaoMarketPictureGetuserpicturesRequest 从 sync.Pool 获取 TaobaoMarketPictureGetuserpicturesAPIRequest
func GetTaobaoMarketPictureGetuserpicturesAPIRequest() *TaobaoMarketPictureGetuserpicturesAPIRequest {
	return poolTaobaoMarketPictureGetuserpicturesAPIRequest.Get().(*TaobaoMarketPictureGetuserpicturesAPIRequest)
}

// ReleaseTaobaoMarketPictureGetuserpicturesAPIRequest 将 TaobaoMarketPictureGetuserpicturesAPIRequest 放入 sync.Pool
func ReleaseTaobaoMarketPictureGetuserpicturesAPIRequest(v *TaobaoMarketPictureGetuserpicturesAPIRequest) {
	v.Reset()
	poolTaobaoMarketPictureGetuserpicturesAPIRequest.Put(v)
}
