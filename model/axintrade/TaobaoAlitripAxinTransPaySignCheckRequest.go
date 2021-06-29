package axintrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿信支付宝验签服务 APIRequest
taobao.alitrip.axin.trans.pay.sign.check

阿信支付宝验签服务
*/
type TaobaoAlitripAxinTransPaySignCheckRequest struct {
    model.Params

    // 验签对象
    axinPayCheckSignDto   *AxinPayCheckSignDto 

}

func NewTaobaoAlitripAxinTransPaySignCheckRequest() *TaobaoAlitripAxinTransPaySignCheckRequest{
    return &TaobaoAlitripAxinTransPaySignCheckRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAlitripAxinTransPaySignCheckRequest) GetApiMethodName() string {
    return "taobao.alitrip.axin.trans.pay.sign.check"
}

func (r TaobaoAlitripAxinTransPaySignCheckRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAlitripAxinTransPaySignCheckRequest) SetAxinPayCheckSignDto(axinPayCheckSignDto *AxinPayCheckSignDto) error {
    r.axinPayCheckSignDto = axinPayCheckSignDto
    r.Set("axin_pay_check_sign_dto", axinPayCheckSignDto)
    return nil
}

func (r TaobaoAlitripAxinTransPaySignCheckRequest) GetAxinPayCheckSignDto() *AxinPayCheckSignDto {
    return r.axinPayCheckSignDto
}

