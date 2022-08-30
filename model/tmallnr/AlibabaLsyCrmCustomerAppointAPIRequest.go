package tmallnr

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLsyCrmCustomerAppointAPIRequest) GetApiMethodName() string {
	return "alibaba.lsy.crm.customer.appoint"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLsyCrmCustomerAppointAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
