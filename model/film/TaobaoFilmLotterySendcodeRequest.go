package film

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘票票外部直发券 APIRequest
taobao.film.lottery.sendcode

淘票票外部直发券
*/
type TaobaoFilmLotterySendcodeRequest struct {
    model.Params

    // 外部商户发码请求
    paramFCodeMerchantSendCodeRequest   *FCodeMerchantSendCodeRq 

}

func NewTaobaoFilmLotterySendcodeRequest() *TaobaoFilmLotterySendcodeRequest{
    return &TaobaoFilmLotterySendcodeRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFilmLotterySendcodeRequest) GetApiMethodName() string {
    return "taobao.film.lottery.sendcode"
}

func (r TaobaoFilmLotterySendcodeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFilmLotterySendcodeRequest) SetParamFCodeMerchantSendCodeRequest(paramFCodeMerchantSendCodeRequest *FCodeMerchantSendCodeRq) error {
    r.paramFCodeMerchantSendCodeRequest = paramFCodeMerchantSendCodeRequest
    r.Set("param_f_code_merchant_send_code_request", paramFCodeMerchantSendCodeRequest)
    return nil
}

func (r TaobaoFilmLotterySendcodeRequest) GetParamFCodeMerchantSendCodeRequest() *FCodeMerchantSendCodeRq {
    return r.paramFCodeMerchantSendCodeRequest
}

