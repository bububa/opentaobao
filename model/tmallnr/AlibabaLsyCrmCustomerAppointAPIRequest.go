package tmallnr

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalsycrmcustomerappointAPIRequest ISV直播间预约 API请求
// alibaba.lsy.crm.customer.appoint
//
// ISV直播间预约
type AlibabalsycrmcustomerappointAPIRequest struct {
	model.Params
	// 入参
	_crmAppointActivityReq *CrmAppointActivityReq
}

// NewAlibabalsycrmcustomerappointRequest 初始化AlibabalsycrmcustomerappointAPIRequest对象
func NewAlibabalsycrmcustomerappointRequest() *AlibabalsycrmcustomerappointAPIRequest {
	return &AlibabalsycrmcustomerappointAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalsycrmcustomerappointAPIRequest) GetApiMethodName() string {
	return "alibaba.lsy.crm.customer.appoint"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalsycrmcustomerappointAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalsycrmcustomerappointAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCrmAppointActivityReq is CrmAppointActivityReq Setter
// 入参
func (r *AlibabalsycrmcustomerappointAPIRequest) SetCrmAppointActivityReq(_crmAppointActivityReq *CrmAppointActivityReq) error {
	r._crmAppointActivityReq = _crmAppointActivityReq
	r.Set("crm_appoint_activity_req", _crmAppointActivityReq)
	return nil
}

// GetCrmAppointActivityReq CrmAppointActivityReq Getter
func (r AlibabalsycrmcustomerappointAPIRequest) GetCrmAppointActivityReq() *CrmAppointActivityReq {
	return r._crmAppointActivityReq
}
