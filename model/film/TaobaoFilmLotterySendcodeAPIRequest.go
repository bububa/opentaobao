package film

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFilmLotterySendcodeAPIRequest 淘票票外部直发券 API请求
// taobao.film.lottery.sendcode
//
// 淘票票外部直发券
type TaobaoFilmLotterySendcodeAPIRequest struct {
	model.Params
	// 外部商户发码请求
	_paramFCodeMerchantSendCodeRequest *FCodeMerchantSendCodeRq
}

// NewTaobaoFilmLotterySendcodeRequest 初始化TaobaoFilmLotterySendcodeAPIRequest对象
func NewTaobaoFilmLotterySendcodeRequest() *TaobaoFilmLotterySendcodeAPIRequest {
	return &TaobaoFilmLotterySendcodeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFilmLotterySendcodeAPIRequest) GetApiMethodName() string {
	return "taobao.film.lottery.sendcode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFilmLotterySendcodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFilmLotterySendcodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamFCodeMerchantSendCodeRequest is ParamFCodeMerchantSendCodeRequest Setter
// 外部商户发码请求
func (r *TaobaoFilmLotterySendcodeAPIRequest) SetParamFCodeMerchantSendCodeRequest(_paramFCodeMerchantSendCodeRequest *FCodeMerchantSendCodeRq) error {
	r._paramFCodeMerchantSendCodeRequest = _paramFCodeMerchantSendCodeRequest
	r.Set("param_f_code_merchant_send_code_request", _paramFCodeMerchantSendCodeRequest)
	return nil
}

// GetParamFCodeMerchantSendCodeRequest ParamFCodeMerchantSendCodeRequest Getter
func (r TaobaoFilmLotterySendcodeAPIRequest) GetParamFCodeMerchantSendCodeRequest() *FCodeMerchantSendCodeRq {
	return r._paramFCodeMerchantSendCodeRequest
}
