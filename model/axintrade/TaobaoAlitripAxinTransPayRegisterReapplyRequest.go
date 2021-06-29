package axintrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿信支付入驻重新申请 APIRequest
taobao.alitrip.axin.trans.pay.register.reapply

阿信支付入驻重新申请
用于支付平台驳回，商户提交时的业务场景
*/
type TaobaoAlitripAxinTransPayRegisterReapplyRequest struct {
    model.Params

    // 阿信支付入驻重新申请参数
    axinPayRegisterCreateDTO   *AxinPayRegisterCreateDto 

}

func NewTaobaoAlitripAxinTransPayRegisterReapplyRequest() *TaobaoAlitripAxinTransPayRegisterReapplyRequest{
    return &TaobaoAlitripAxinTransPayRegisterReapplyRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAlitripAxinTransPayRegisterReapplyRequest) GetApiMethodName() string {
    return "taobao.alitrip.axin.trans.pay.register.reapply"
}

func (r TaobaoAlitripAxinTransPayRegisterReapplyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAlitripAxinTransPayRegisterReapplyRequest) SetAxinPayRegisterCreateDTO(axinPayRegisterCreateDTO *AxinPayRegisterCreateDto) error {
    r.axinPayRegisterCreateDTO = axinPayRegisterCreateDTO
    r.Set("axin_pay_register_create_d_t_o", axinPayRegisterCreateDTO)
    return nil
}

func (r TaobaoAlitripAxinTransPayRegisterReapplyRequest) GetAxinPayRegisterCreateDTO() *AxinPayRegisterCreateDto {
    return r.axinPayRegisterCreateDTO
}

