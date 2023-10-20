package tmallnr

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLsyCrmCustomerAppointAPIRequest ISV直播间预约 API请求
// alibaba.lsy.crm.customer.appoint
//
// ISV直播间预约
type AlibabaLsyCrmCustomerAppointAPIRequest struct {
	model.Params
	// 入参
	_crmAppointActivityReq *CrmAppointActivityReq
}

// NewAlibabaLsyCrmCustomerAppointRequest 初始化AlibabaLsyCrmCustomerAppointAPIRequest对象
func NewAlibabaLsyCrmCustomerAppointRequest() *AlibabaLsyCrmCustomerAppointAPIRequest {
	return &AlibabaLsyCrmCustomerAppointAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLsyCrmCustomerAppointAPIRequest) Reset() {
	r._crmAppointActivityReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLsyCrmCustomerAppointAPIRequest) GetApiMethodName() string {
	return "alibaba.lsy.crm.customer.appoint"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLsyCrmCustomerAppointAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLsyCrmCustomerAppointAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCrmAppointActivityReq is CrmAppointActivityReq Setter
// 入参
func (r *AlibabaLsyCrmCustomerAppointAPIRequest) SetCrmAppointActivityReq(_crmAppointActivityReq *CrmAppointActivityReq) error {
	r._crmAppointActivityReq = _crmAppointActivityReq
	r.Set("crm_appoint_activity_req", _crmAppointActivityReq)
	return nil
}

// GetCrmAppointActivityReq CrmAppointActivityReq Getter
func (r AlibabaLsyCrmCustomerAppointAPIRequest) GetCrmAppointActivityReq() *CrmAppointActivityReq {
	return r._crmAppointActivityReq
}

var poolAlibabaLsyCrmCustomerAppointAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLsyCrmCustomerAppointRequest()
	},
}

// GetAlibabaLsyCrmCustomerAppointRequest 从 sync.Pool 获取 AlibabaLsyCrmCustomerAppointAPIRequest
func GetAlibabaLsyCrmCustomerAppointAPIRequest() *AlibabaLsyCrmCustomerAppointAPIRequest {
	return poolAlibabaLsyCrmCustomerAppointAPIRequest.Get().(*AlibabaLsyCrmCustomerAppointAPIRequest)
}

// ReleaseAlibabaLsyCrmCustomerAppointAPIRequest 将 AlibabaLsyCrmCustomerAppointAPIRequest 放入 sync.Pool
func ReleaseAlibabaLsyCrmCustomerAppointAPIRequest(v *AlibabaLsyCrmCustomerAppointAPIRequest) {
	v.Reset()
	poolAlibabaLsyCrmCustomerAppointAPIRequest.Put(v)
}
