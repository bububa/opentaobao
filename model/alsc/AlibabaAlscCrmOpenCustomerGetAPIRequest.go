package alsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmOpenCustomerGetAPIRequest 查询会员资产 API请求
// alibaba.alsc.crm.open.customer.get
//
// 查询会员身份，资产等
type AlibabaAlscCrmOpenCustomerGetAPIRequest struct {
	model.Params
	// 入参
	_paramCustomerGetOpenReq *CustomerGetOpenReq
}

// NewAlibabaAlscCrmOpenCustomerGetRequest 初始化AlibabaAlscCrmOpenCustomerGetAPIRequest对象
func NewAlibabaAlscCrmOpenCustomerGetRequest() *AlibabaAlscCrmOpenCustomerGetAPIRequest {
	return &AlibabaAlscCrmOpenCustomerGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlscCrmOpenCustomerGetAPIRequest) Reset() {
	r._paramCustomerGetOpenReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmOpenCustomerGetAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.open.customer.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmOpenCustomerGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscCrmOpenCustomerGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamCustomerGetOpenReq is ParamCustomerGetOpenReq Setter
// 入参
func (r *AlibabaAlscCrmOpenCustomerGetAPIRequest) SetParamCustomerGetOpenReq(_paramCustomerGetOpenReq *CustomerGetOpenReq) error {
	r._paramCustomerGetOpenReq = _paramCustomerGetOpenReq
	r.Set("param_customer_get_open_req", _paramCustomerGetOpenReq)
	return nil
}

// GetParamCustomerGetOpenReq ParamCustomerGetOpenReq Getter
func (r AlibabaAlscCrmOpenCustomerGetAPIRequest) GetParamCustomerGetOpenReq() *CustomerGetOpenReq {
	return r._paramCustomerGetOpenReq
}

var poolAlibabaAlscCrmOpenCustomerGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlscCrmOpenCustomerGetRequest()
	},
}

// GetAlibabaAlscCrmOpenCustomerGetRequest 从 sync.Pool 获取 AlibabaAlscCrmOpenCustomerGetAPIRequest
func GetAlibabaAlscCrmOpenCustomerGetAPIRequest() *AlibabaAlscCrmOpenCustomerGetAPIRequest {
	return poolAlibabaAlscCrmOpenCustomerGetAPIRequest.Get().(*AlibabaAlscCrmOpenCustomerGetAPIRequest)
}

// ReleaseAlibabaAlscCrmOpenCustomerGetAPIRequest 将 AlibabaAlscCrmOpenCustomerGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlscCrmOpenCustomerGetAPIRequest(v *AlibabaAlscCrmOpenCustomerGetAPIRequest) {
	v.Reset()
	poolAlibabaAlscCrmOpenCustomerGetAPIRequest.Put(v)
}
