package alsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmCustomerCreateAPIRequest 创建顾客 API请求
// alibaba.alsc.crm.customer.create
//
// 开放本地生活创建顾客功能
type AlibabaAlscCrmCustomerCreateAPIRequest struct {
	model.Params
	// 创建顾客参数
	_paramCustomerCreateOpenReq *CustomerCreateOpenReq
}

// NewAlibabaAlscCrmCustomerCreateRequest 初始化AlibabaAlscCrmCustomerCreateAPIRequest对象
func NewAlibabaAlscCrmCustomerCreateRequest() *AlibabaAlscCrmCustomerCreateAPIRequest {
	return &AlibabaAlscCrmCustomerCreateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlscCrmCustomerCreateAPIRequest) Reset() {
	r._paramCustomerCreateOpenReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmCustomerCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.customer.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmCustomerCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscCrmCustomerCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamCustomerCreateOpenReq is ParamCustomerCreateOpenReq Setter
// 创建顾客参数
func (r *AlibabaAlscCrmCustomerCreateAPIRequest) SetParamCustomerCreateOpenReq(_paramCustomerCreateOpenReq *CustomerCreateOpenReq) error {
	r._paramCustomerCreateOpenReq = _paramCustomerCreateOpenReq
	r.Set("param_customer_create_open_req", _paramCustomerCreateOpenReq)
	return nil
}

// GetParamCustomerCreateOpenReq ParamCustomerCreateOpenReq Getter
func (r AlibabaAlscCrmCustomerCreateAPIRequest) GetParamCustomerCreateOpenReq() *CustomerCreateOpenReq {
	return r._paramCustomerCreateOpenReq
}

var poolAlibabaAlscCrmCustomerCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlscCrmCustomerCreateRequest()
	},
}

// GetAlibabaAlscCrmCustomerCreateRequest 从 sync.Pool 获取 AlibabaAlscCrmCustomerCreateAPIRequest
func GetAlibabaAlscCrmCustomerCreateAPIRequest() *AlibabaAlscCrmCustomerCreateAPIRequest {
	return poolAlibabaAlscCrmCustomerCreateAPIRequest.Get().(*AlibabaAlscCrmCustomerCreateAPIRequest)
}

// ReleaseAlibabaAlscCrmCustomerCreateAPIRequest 将 AlibabaAlscCrmCustomerCreateAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlscCrmCustomerCreateAPIRequest(v *AlibabaAlscCrmCustomerCreateAPIRequest) {
	v.Reset()
	poolAlibabaAlscCrmCustomerCreateAPIRequest.Put(v)
}
