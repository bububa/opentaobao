package film

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘票票外部直发券 API请求
taobao.film.lottery.sendcode

淘票票外部直发券
*/
type TaobaoFilmLotterySendcodeRequest struct {
    model.Params
    // 外部商户发码请求
    _paramFCodeMerchantSendCodeRequest   *FCodeMerchantSendCodeRq
}

// 初始化TaobaoFilmLotterySendcodeRequest对象
func NewTaobaoFilmLotterySendcodeRequest() *TaobaoFilmLotterySendcodeRequest{
    return &TaobaoFilmLotterySendcodeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFilmLotterySendcodeRequest) GetApiMethodName() string {
    return "taobao.film.lottery.sendcode"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFilmLotterySendcodeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamFCodeMerchantSendCodeRequest Setter
// 外部商户发码请求
func (r *TaobaoFilmLotterySendcodeRequest) SetParamFCodeMerchantSendCodeRequest(_paramFCodeMerchantSendCodeRequest *FCodeMerchantSendCodeRq) error {
    r._paramFCodeMerchantSendCodeRequest = _paramFCodeMerchantSendCodeRequest
    r.Set("param_f_code_merchant_send_code_request", _paramFCodeMerchantSendCodeRequest)
    return nil
}

// ParamFCodeMerchantSendCodeRequest Getter
func (r TaobaoFilmLotterySendcodeRequest) GetParamFCodeMerchantSendCodeRequest() *FCodeMerchantSendCodeRq {
    return r._paramFCodeMerchantSendCodeRequest
}
