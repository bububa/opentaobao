package servicecenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
内购服务确认单申请接口 APIRequest
taobao.fuwu.sp.confirm.apply

isv能通过该接口发起确认申请单
*/
type TaobaoFuwuSpConfirmApplyRequest struct {
    model.Params

    // 确认单申请
    paramIncomeConfirmDTO   *IncomeConfirmDto 

}

func NewTaobaoFuwuSpConfirmApplyRequest() *TaobaoFuwuSpConfirmApplyRequest{
    return &TaobaoFuwuSpConfirmApplyRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFuwuSpConfirmApplyRequest) GetApiMethodName() string {
    return "taobao.fuwu.sp.confirm.apply"
}

func (r TaobaoFuwuSpConfirmApplyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFuwuSpConfirmApplyRequest) SetParamIncomeConfirmDTO(paramIncomeConfirmDTO *IncomeConfirmDto) error {
    r.paramIncomeConfirmDTO = paramIncomeConfirmDTO
    r.Set("param_income_confirm_d_t_o", paramIncomeConfirmDTO)
    return nil
}

func (r TaobaoFuwuSpConfirmApplyRequest) GetParamIncomeConfirmDTO() *IncomeConfirmDto {
    return r.paramIncomeConfirmDTO
}

