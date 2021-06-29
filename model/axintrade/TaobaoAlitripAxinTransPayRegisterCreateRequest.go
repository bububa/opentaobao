package axintrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
提交支付服务开通 APIRequest
taobao.alitrip.axin.trans.pay.register.create

阿信供销平台-提交支付服务开通接口
*/
type TaobaoAlitripAxinTransPayRegisterCreateRequest struct {
    model.Params

    // 提交支付服务开通接口入参
    createDTO   *AxinPayRegisterCreateDto 

}

func NewTaobaoAlitripAxinTransPayRegisterCreateRequest() *TaobaoAlitripAxinTransPayRegisterCreateRequest{
    return &TaobaoAlitripAxinTransPayRegisterCreateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAlitripAxinTransPayRegisterCreateRequest) GetApiMethodName() string {
    return "taobao.alitrip.axin.trans.pay.register.create"
}

func (r TaobaoAlitripAxinTransPayRegisterCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAlitripAxinTransPayRegisterCreateRequest) SetCreateDTO(createDTO *AxinPayRegisterCreateDto) error {
    r.createDTO = createDTO
    r.Set("create_d_t_o", createDTO)
    return nil
}

func (r TaobaoAlitripAxinTransPayRegisterCreateRequest) GetCreateDTO() *AxinPayRegisterCreateDto {
    return r.createDTO
}

