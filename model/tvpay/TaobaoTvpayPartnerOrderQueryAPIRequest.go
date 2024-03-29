package tvpay

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTvpayPartnerOrderQueryAPIRequest 商户查询订单 API请求
// taobao.tvpay.partner.order.query
//
// 给商户提供的查询订单状态的API
type TaobaoTvpayPartnerOrderQueryAPIRequest struct {
	model.Params
	// 商户订单号
	_orderNo string
}

// NewTaobaoTvpayPartnerOrderQueryRequest 初始化TaobaoTvpayPartnerOrderQueryAPIRequest对象
func NewTaobaoTvpayPartnerOrderQueryRequest() *TaobaoTvpayPartnerOrderQueryAPIRequest {
	return &TaobaoTvpayPartnerOrderQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTvpayPartnerOrderQueryAPIRequest) Reset() {
	r._orderNo = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTvpayPartnerOrderQueryAPIRequest) GetApiMethodName() string {
	return "taobao.tvpay.partner.order.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTvpayPartnerOrderQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTvpayPartnerOrderQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderNo is OrderNo Setter
// 商户订单号
func (r *TaobaoTvpayPartnerOrderQueryAPIRequest) SetOrderNo(_orderNo string) error {
	r._orderNo = _orderNo
	r.Set("order_no", _orderNo)
	return nil
}

// GetOrderNo OrderNo Getter
func (r TaobaoTvpayPartnerOrderQueryAPIRequest) GetOrderNo() string {
	return r._orderNo
}

var poolTaobaoTvpayPartnerOrderQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTvpayPartnerOrderQueryRequest()
	},
}

// GetTaobaoTvpayPartnerOrderQueryRequest 从 sync.Pool 获取 TaobaoTvpayPartnerOrderQueryAPIRequest
func GetTaobaoTvpayPartnerOrderQueryAPIRequest() *TaobaoTvpayPartnerOrderQueryAPIRequest {
	return poolTaobaoTvpayPartnerOrderQueryAPIRequest.Get().(*TaobaoTvpayPartnerOrderQueryAPIRequest)
}

// ReleaseTaobaoTvpayPartnerOrderQueryAPIRequest 将 TaobaoTvpayPartnerOrderQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoTvpayPartnerOrderQueryAPIRequest(v *TaobaoTvpayPartnerOrderQueryAPIRequest) {
	v.Reset()
	poolTaobaoTvpayPartnerOrderQueryAPIRequest.Put(v)
}
