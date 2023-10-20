package film

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofilmlotterysendcodeAPIRequest 淘票票外部直发券 API请求
// taobao.film.lottery.sendcode
//
// 淘票票外部直发券
type TaobaofilmlotterysendcodeAPIRequest struct {
	model.Params
	// 外部商户发码请求
	_paramFCodeMerchantSendCodeRequest *FcodeMerchantSendCodeRq
}

// NewTaobaofilmlotterysendcodeRequest 初始化TaobaofilmlotterysendcodeAPIRequest对象
func NewTaobaofilmlotterysendcodeRequest() *TaobaofilmlotterysendcodeAPIRequest {
	return &TaobaofilmlotterysendcodeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofilmlotterysendcodeAPIRequest) GetApiMethodName() string {
	return "taobao.film.lottery.sendcode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofilmlotterysendcodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofilmlotterysendcodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamFCodeMerchantSendCodeRequest is ParamFCodeMerchantSendCodeRequest Setter
// 外部商户发码请求
func (r *TaobaofilmlotterysendcodeAPIRequest) SetParamFCodeMerchantSendCodeRequest(_paramFCodeMerchantSendCodeRequest *FcodeMerchantSendCodeRq) error {
	r._paramFCodeMerchantSendCodeRequest = _paramFCodeMerchantSendCodeRequest
	r.Set("param_f_code_merchant_send_code_request", _paramFCodeMerchantSendCodeRequest)
	return nil
}

// GetParamFCodeMerchantSendCodeRequest ParamFCodeMerchantSendCodeRequest Getter
func (r TaobaofilmlotterysendcodeAPIRequest) GetParamFCodeMerchantSendCodeRequest() *FcodeMerchantSendCodeRq {
	return r._paramFCodeMerchantSendCodeRequest
}
