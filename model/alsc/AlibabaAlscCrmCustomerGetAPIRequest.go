package alsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmCustomerGetAPIRequest 查询顾客详情 API请求
// alibaba.alsc.crm.customer.get
//
// 查询顾客详情
type AlibabaAlscCrmCustomerGetAPIRequest struct {
	model.Params
	// 顾客详情查询条件
	_paramCustomerIdQueryOpenReq *CustomerIdQueryOpenReq
}

// NewAlibabaAlscCrmCustomerGetRequest 初始化AlibabaAlscCrmCustomerGetAPIRequest对象
func NewAlibabaAlscCrmCustomerGetRequest() *AlibabaAlscCrmCustomerGetAPIRequest {
	return &AlibabaAlscCrmCustomerGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlscCrmCustomerGetAPIRequest) Reset() {
	r._paramCustomerIdQueryOpenReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmCustomerGetAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.customer.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmCustomerGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscCrmCustomerGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamCustomerIdQueryOpenReq is ParamCustomerIdQueryOpenReq Setter
// 顾客详情查询条件
func (r *AlibabaAlscCrmCustomerGetAPIRequest) SetParamCustomerIdQueryOpenReq(_paramCustomerIdQueryOpenReq *CustomerIdQueryOpenReq) error {
	r._paramCustomerIdQueryOpenReq = _paramCustomerIdQueryOpenReq
	r.Set("param_customer_id_query_open_req", _paramCustomerIdQueryOpenReq)
	return nil
}

// GetParamCustomerIdQueryOpenReq ParamCustomerIdQueryOpenReq Getter
func (r AlibabaAlscCrmCustomerGetAPIRequest) GetParamCustomerIdQueryOpenReq() *CustomerIdQueryOpenReq {
	return r._paramCustomerIdQueryOpenReq
}

var poolAlibabaAlscCrmCustomerGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlscCrmCustomerGetRequest()
	},
}

// GetAlibabaAlscCrmCustomerGetRequest 从 sync.Pool 获取 AlibabaAlscCrmCustomerGetAPIRequest
func GetAlibabaAlscCrmCustomerGetAPIRequest() *AlibabaAlscCrmCustomerGetAPIRequest {
	return poolAlibabaAlscCrmCustomerGetAPIRequest.Get().(*AlibabaAlscCrmCustomerGetAPIRequest)
}

// ReleaseAlibabaAlscCrmCustomerGetAPIRequest 将 AlibabaAlscCrmCustomerGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlscCrmCustomerGetAPIRequest(v *AlibabaAlscCrmCustomerGetAPIRequest) {
	v.Reset()
	poolAlibabaAlscCrmCustomerGetAPIRequest.Put(v)
}
