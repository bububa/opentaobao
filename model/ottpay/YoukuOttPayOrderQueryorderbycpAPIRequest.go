package ottpay

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YoukuOttPayOrderQueryorderbycpAPIRequest
订单查询接口(cp订单号查询) API请求
youku.ott.pay.order.queryorderbycp

给商户服务端查询订单状态 */
type YoukuOttPayOrderQueryorderbycpAPIRequest struct {
	model.Params
	// cp订单号
	_cpOrderNo string
}

// NewYoukuOttPayOrderQueryorderbycpRequest 初始化YoukuOttPayOrderQueryorderbycpAPIRequest对象
func NewYoukuOttPayOrderQueryorderbycpRequest() *YoukuOttPayOrderQueryorderbycpAPIRequest {
	return &YoukuOttPayOrderQueryorderbycpAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YoukuOttPayOrderQueryorderbycpAPIRequest) GetApiMethodName() string {
	return "youku.ott.pay.order.queryorderbycp"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YoukuOttPayOrderQueryorderbycpAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is CpOrderNo Setter
// cp订单号
func (r *YoukuOttPayOrderQueryorderbycpAPIRequest) SetCpOrderNo(_cpOrderNo string) error {
	r._cpOrderNo = _cpOrderNo
	r.Set("cp_order_no", _cpOrderNo)
	return nil
}

// Get CpOrderNo Getter
func (r YoukuOttPayOrderQueryorderbycpAPIRequest) GetCpOrderNo() string {
	return r._cpOrderNo
}
