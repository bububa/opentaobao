package alsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmCardBindcustomerAPIRequest 卡号绑定顾客 API请求
// alibaba.alsc.crm.card.bindcustomer
//
// 为卡号绑定顾客
type AlibabaAlscCrmCardBindcustomerAPIRequest struct {
	model.Params
	// 请求参数
	_paramBindCustomerOpenReq *BindCustomerOpenReq
}

// NewAlibabaAlscCrmCardBindcustomerRequest 初始化AlibabaAlscCrmCardBindcustomerAPIRequest对象
func NewAlibabaAlscCrmCardBindcustomerRequest() *AlibabaAlscCrmCardBindcustomerAPIRequest {
	return &AlibabaAlscCrmCardBindcustomerAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlscCrmCardBindcustomerAPIRequest) Reset() {
	r._paramBindCustomerOpenReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmCardBindcustomerAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.card.bindcustomer"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmCardBindcustomerAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscCrmCardBindcustomerAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBindCustomerOpenReq is ParamBindCustomerOpenReq Setter
// 请求参数
func (r *AlibabaAlscCrmCardBindcustomerAPIRequest) SetParamBindCustomerOpenReq(_paramBindCustomerOpenReq *BindCustomerOpenReq) error {
	r._paramBindCustomerOpenReq = _paramBindCustomerOpenReq
	r.Set("param_bind_customer_open_req", _paramBindCustomerOpenReq)
	return nil
}

// GetParamBindCustomerOpenReq ParamBindCustomerOpenReq Getter
func (r AlibabaAlscCrmCardBindcustomerAPIRequest) GetParamBindCustomerOpenReq() *BindCustomerOpenReq {
	return r._paramBindCustomerOpenReq
}

var poolAlibabaAlscCrmCardBindcustomerAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlscCrmCardBindcustomerRequest()
	},
}

// GetAlibabaAlscCrmCardBindcustomerRequest 从 sync.Pool 获取 AlibabaAlscCrmCardBindcustomerAPIRequest
func GetAlibabaAlscCrmCardBindcustomerAPIRequest() *AlibabaAlscCrmCardBindcustomerAPIRequest {
	return poolAlibabaAlscCrmCardBindcustomerAPIRequest.Get().(*AlibabaAlscCrmCardBindcustomerAPIRequest)
}

// ReleaseAlibabaAlscCrmCardBindcustomerAPIRequest 将 AlibabaAlscCrmCardBindcustomerAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlscCrmCardBindcustomerAPIRequest(v *AlibabaAlscCrmCardBindcustomerAPIRequest) {
	v.Reset()
	poolAlibabaAlscCrmCardBindcustomerAPIRequest.Put(v)
}
