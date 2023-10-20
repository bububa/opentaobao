package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofuwuspconfirmapplyAPIRequest 内购服务确认单申请接口 API请求
// taobao.fuwu.sp.confirm.apply
//
// isv能通过该接口发起确认申请单
type TaobaofuwuspconfirmapplyAPIRequest struct {
	model.Params
	// 确认单申请
	_paramIncomeConfirmDTO *IncomeConfirmDto
}

// NewTaobaofuwuspconfirmapplyRequest 初始化TaobaofuwuspconfirmapplyAPIRequest对象
func NewTaobaofuwuspconfirmapplyRequest() *TaobaofuwuspconfirmapplyAPIRequest {
	return &TaobaofuwuspconfirmapplyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofuwuspconfirmapplyAPIRequest) GetApiMethodName() string {
	return "taobao.fuwu.sp.confirm.apply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofuwuspconfirmapplyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofuwuspconfirmapplyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamIncomeConfirmDTO is ParamIncomeConfirmDTO Setter
// 确认单申请
func (r *TaobaofuwuspconfirmapplyAPIRequest) SetParamIncomeConfirmDTO(_paramIncomeConfirmDTO *IncomeConfirmDto) error {
	r._paramIncomeConfirmDTO = _paramIncomeConfirmDTO
	r.Set("param_income_confirm_d_t_o", _paramIncomeConfirmDTO)
	return nil
}

// GetParamIncomeConfirmDTO ParamIncomeConfirmDTO Getter
func (r TaobaofuwuspconfirmapplyAPIRequest) GetParamIncomeConfirmDTO() *IncomeConfirmDto {
	return r._paramIncomeConfirmDTO
}
