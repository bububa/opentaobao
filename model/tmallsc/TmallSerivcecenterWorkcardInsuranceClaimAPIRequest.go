package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallserivcecenterworkcardinsuranceclaimAPIRequest 保险理赔回传工单记录 API请求
// tmall.serivcecenter.workcard.insurance.claim
//
// 保险理赔回传工单记录
type TmallserivcecenterworkcardinsuranceclaimAPIRequest struct {
	model.Params
	// 工单回传理赔信息入参
	_workcardInsuranceCallbackRequest *WorkcardInsuranceCallbackRequest
}

// NewTmallserivcecenterworkcardinsuranceclaimRequest 初始化TmallserivcecenterworkcardinsuranceclaimAPIRequest对象
func NewTmallserivcecenterworkcardinsuranceclaimRequest() *TmallserivcecenterworkcardinsuranceclaimAPIRequest {
	return &TmallserivcecenterworkcardinsuranceclaimAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallserivcecenterworkcardinsuranceclaimAPIRequest) GetApiMethodName() string {
	return "tmall.serivcecenter.workcard.insurance.claim"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallserivcecenterworkcardinsuranceclaimAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallserivcecenterworkcardinsuranceclaimAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWorkcardInsuranceCallbackRequest is WorkcardInsuranceCallbackRequest Setter
// 工单回传理赔信息入参
func (r *TmallserivcecenterworkcardinsuranceclaimAPIRequest) SetWorkcardInsuranceCallbackRequest(_workcardInsuranceCallbackRequest *WorkcardInsuranceCallbackRequest) error {
	r._workcardInsuranceCallbackRequest = _workcardInsuranceCallbackRequest
	r.Set("workcard_insurance_callback_request", _workcardInsuranceCallbackRequest)
	return nil
}

// GetWorkcardInsuranceCallbackRequest WorkcardInsuranceCallbackRequest Getter
func (r TmallserivcecenterworkcardinsuranceclaimAPIRequest) GetWorkcardInsuranceCallbackRequest() *WorkcardInsuranceCallbackRequest {
	return r._workcardInsuranceCallbackRequest
}
