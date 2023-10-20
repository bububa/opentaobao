package ottpay

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YoukuottpayorderqueryorderbycpAPIRequest 订单查询接口(cp订单号查询) API请求
// youku.ott.pay.order.queryorderbycp
//
// 给商户服务端查询订单状态
type YoukuottpayorderqueryorderbycpAPIRequest struct {
	model.Params
	// cp订单号
	_cpOrderNo string
}

// NewYoukuottpayorderqueryorderbycpRequest 初始化YoukuottpayorderqueryorderbycpAPIRequest对象
func NewYoukuottpayorderqueryorderbycpRequest() *YoukuottpayorderqueryorderbycpAPIRequest {
	return &YoukuottpayorderqueryorderbycpAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YoukuottpayorderqueryorderbycpAPIRequest) GetApiMethodName() string {
	return "youku.ott.pay.order.queryorderbycp"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YoukuottpayorderqueryorderbycpAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YoukuottpayorderqueryorderbycpAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCpOrderNo is CpOrderNo Setter
// cp订单号
func (r *YoukuottpayorderqueryorderbycpAPIRequest) SetCpOrderNo(_cpOrderNo string) error {
	r._cpOrderNo = _cpOrderNo
	r.Set("cp_order_no", _cpOrderNo)
	return nil
}

// GetCpOrderNo CpOrderNo Getter
func (r YoukuottpayorderqueryorderbycpAPIRequest) GetCpOrderNo() string {
	return r._cpOrderNo
}
