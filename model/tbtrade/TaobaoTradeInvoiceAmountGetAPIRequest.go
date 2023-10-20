package tbtrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotradeinvoiceamountgetAPIRequest 获取订单应开票金额 API请求
// taobao.trade.invoice.amount.get
//
// 订单应开票金额计算
type TaobaotradeinvoiceamountgetAPIRequest struct {
	model.Params
	// 业务订单ID
	_tid int64
}

// NewTaobaotradeinvoiceamountgetRequest 初始化TaobaotradeinvoiceamountgetAPIRequest对象
func NewTaobaotradeinvoiceamountgetRequest() *TaobaotradeinvoiceamountgetAPIRequest {
	return &TaobaotradeinvoiceamountgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotradeinvoiceamountgetAPIRequest) GetApiMethodName() string {
	return "taobao.trade.invoice.amount.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotradeinvoiceamountgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotradeinvoiceamountgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTid is Tid Setter
// 业务订单ID
func (r *TaobaotradeinvoiceamountgetAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaotradeinvoiceamountgetAPIRequest) GetTid() int64 {
	return r._tid
}
