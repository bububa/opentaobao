package ottpay

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YoukuOttPayOrderQueryorderbycpAPIRequest 订单查询接口(cp订单号查询) API请求
// youku.ott.pay.order.queryorderbycp
//
// 给商户服务端查询订单状态
type YoukuOttPayOrderQueryorderbycpAPIRequest struct {
	model.Params
	// cp订单号
	_cpOrderNo string
}

// NewYoukuOttPayOrderQueryorderbycpRequest 初始化YoukuOttPayOrderQueryorderbycpAPIRequest对象
func NewYoukuOttPayOrderQueryorderbycpRequest() *YoukuOttPayOrderQueryorderbycpAPIRequest {
	return &YoukuOttPayOrderQueryorderbycpAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YoukuOttPayOrderQueryorderbycpAPIRequest) Reset() {
	r._cpOrderNo = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YoukuOttPayOrderQueryorderbycpAPIRequest) GetApiMethodName() string {
	return "youku.ott.pay.order.queryorderbycp"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YoukuOttPayOrderQueryorderbycpAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YoukuOttPayOrderQueryorderbycpAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCpOrderNo is CpOrderNo Setter
// cp订单号
func (r *YoukuOttPayOrderQueryorderbycpAPIRequest) SetCpOrderNo(_cpOrderNo string) error {
	r._cpOrderNo = _cpOrderNo
	r.Set("cp_order_no", _cpOrderNo)
	return nil
}

// GetCpOrderNo CpOrderNo Getter
func (r YoukuOttPayOrderQueryorderbycpAPIRequest) GetCpOrderNo() string {
	return r._cpOrderNo
}

var poolYoukuOttPayOrderQueryorderbycpAPIRequest = sync.Pool{
	New: func() any {
		return NewYoukuOttPayOrderQueryorderbycpRequest()
	},
}

// GetYoukuOttPayOrderQueryorderbycpRequest 从 sync.Pool 获取 YoukuOttPayOrderQueryorderbycpAPIRequest
func GetYoukuOttPayOrderQueryorderbycpAPIRequest() *YoukuOttPayOrderQueryorderbycpAPIRequest {
	return poolYoukuOttPayOrderQueryorderbycpAPIRequest.Get().(*YoukuOttPayOrderQueryorderbycpAPIRequest)
}

// ReleaseYoukuOttPayOrderQueryorderbycpAPIRequest 将 YoukuOttPayOrderQueryorderbycpAPIRequest 放入 sync.Pool
func ReleaseYoukuOttPayOrderQueryorderbycpAPIRequest(v *YoukuOttPayOrderQueryorderbycpAPIRequest) {
	v.Reset()
	poolYoukuOttPayOrderQueryorderbycpAPIRequest.Put(v)
}
