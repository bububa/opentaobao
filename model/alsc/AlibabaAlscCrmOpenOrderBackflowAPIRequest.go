package alsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmOpenOrderBackflowAPIRequest 订单回流接口 API请求
// alibaba.alsc.crm.open.order.backflow
//
// 回流isv订单接口
type AlibabaAlscCrmOpenOrderBackflowAPIRequest struct {
	model.Params
	// 入参
	_paramOrderBackflowOpenReq *OrderBackflowOpenReq
}

// NewAlibabaAlscCrmOpenOrderBackflowRequest 初始化AlibabaAlscCrmOpenOrderBackflowAPIRequest对象
func NewAlibabaAlscCrmOpenOrderBackflowRequest() *AlibabaAlscCrmOpenOrderBackflowAPIRequest {
	return &AlibabaAlscCrmOpenOrderBackflowAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlscCrmOpenOrderBackflowAPIRequest) Reset() {
	r._paramOrderBackflowOpenReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmOpenOrderBackflowAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.open.order.backflow"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmOpenOrderBackflowAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscCrmOpenOrderBackflowAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamOrderBackflowOpenReq is ParamOrderBackflowOpenReq Setter
// 入参
func (r *AlibabaAlscCrmOpenOrderBackflowAPIRequest) SetParamOrderBackflowOpenReq(_paramOrderBackflowOpenReq *OrderBackflowOpenReq) error {
	r._paramOrderBackflowOpenReq = _paramOrderBackflowOpenReq
	r.Set("param_order_backflow_open_req", _paramOrderBackflowOpenReq)
	return nil
}

// GetParamOrderBackflowOpenReq ParamOrderBackflowOpenReq Getter
func (r AlibabaAlscCrmOpenOrderBackflowAPIRequest) GetParamOrderBackflowOpenReq() *OrderBackflowOpenReq {
	return r._paramOrderBackflowOpenReq
}

var poolAlibabaAlscCrmOpenOrderBackflowAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlscCrmOpenOrderBackflowRequest()
	},
}

// GetAlibabaAlscCrmOpenOrderBackflowRequest 从 sync.Pool 获取 AlibabaAlscCrmOpenOrderBackflowAPIRequest
func GetAlibabaAlscCrmOpenOrderBackflowAPIRequest() *AlibabaAlscCrmOpenOrderBackflowAPIRequest {
	return poolAlibabaAlscCrmOpenOrderBackflowAPIRequest.Get().(*AlibabaAlscCrmOpenOrderBackflowAPIRequest)
}

// ReleaseAlibabaAlscCrmOpenOrderBackflowAPIRequest 将 AlibabaAlscCrmOpenOrderBackflowAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlscCrmOpenOrderBackflowAPIRequest(v *AlibabaAlscCrmOpenOrderBackflowAPIRequest) {
	v.Reset()
	poolAlibabaAlscCrmOpenOrderBackflowAPIRequest.Put(v)
}
