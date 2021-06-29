package servicecenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
内购服务确认单申请接口 API请求
taobao.fuwu.sp.confirm.apply

isv能通过该接口发起确认申请单
*/
type TaobaoFuwuSpConfirmApplyRequest struct {
    model.Params
    // 确认单申请
    _paramIncomeConfirmDTO   *IncomeConfirmDTO
}

// 初始化TaobaoFuwuSpConfirmApplyRequest对象
func NewTaobaoFuwuSpConfirmApplyRequest() *TaobaoFuwuSpConfirmApplyRequest{
    return &TaobaoFuwuSpConfirmApplyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFuwuSpConfirmApplyRequest) GetApiMethodName() string {
    return "taobao.fuwu.sp.confirm.apply"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFuwuSpConfirmApplyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamIncomeConfirmDTO Setter
// 确认单申请
func (r *TaobaoFuwuSpConfirmApplyRequest) SetParamIncomeConfirmDTO(_paramIncomeConfirmDTO *IncomeConfirmDTO) error {
    r._paramIncomeConfirmDTO = _paramIncomeConfirmDTO
    r.Set("param_income_confirm_d_t_o", _paramIncomeConfirmDTO)
    return nil
}

// ParamIncomeConfirmDTO Getter
func (r TaobaoFuwuSpConfirmApplyRequest) GetParamIncomeConfirmDTO() *IncomeConfirmDTO {
    return r._paramIncomeConfirmDTO
}
