package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFuwuSpConfirmApplyAPIRequest
内购服务确认单申请接口 API请求
taobao.fuwu.sp.confirm.apply

isv能通过该接口发起确认申请单 */
type TaobaoFuwuSpConfirmApplyAPIRequest struct {
	model.Params
	// 确认单申请
	_paramIncomeConfirmDTO *IncomeConfirmDto
}

// NewTaobaoFuwuSpConfirmApplyRequest 初始化TaobaoFuwuSpConfirmApplyAPIRequest对象
func NewTaobaoFuwuSpConfirmApplyRequest() *TaobaoFuwuSpConfirmApplyAPIRequest {
	return &TaobaoFuwuSpConfirmApplyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFuwuSpConfirmApplyAPIRequest) GetApiMethodName() string {
	return "taobao.fuwu.sp.confirm.apply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFuwuSpConfirmApplyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamIncomeConfirmDTO Setter
// 确认单申请
func (r *TaobaoFuwuSpConfirmApplyAPIRequest) SetParamIncomeConfirmDTO(_paramIncomeConfirmDTO *IncomeConfirmDto) error {
	r._paramIncomeConfirmDTO = _paramIncomeConfirmDTO
	r.Set("param_income_confirm_d_t_o", _paramIncomeConfirmDTO)
	return nil
}

// Get ParamIncomeConfirmDTO Getter
func (r TaobaoFuwuSpConfirmApplyAPIRequest) GetParamIncomeConfirmDTO() *IncomeConfirmDto {
	return r._paramIncomeConfirmDTO
}
