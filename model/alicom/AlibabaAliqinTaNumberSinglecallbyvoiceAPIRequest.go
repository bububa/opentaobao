package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAliqinTaNumberSinglecallbyvoiceAPIRequest
根据号码tts单呼 API请求
alibaba.aliqin.ta.number.singlecallbyvoice

根据号码语音单呼 */
type AlibabaAliqinTaNumberSinglecallbyvoiceAPIRequest struct {
	model.Params
	// 单呼号码
	_calledNum string
	// 显示号码
	_calledShowNum string
	// 语音文件code
	_voiceCode string
	// 上下文参数 示例:{"extend":"回传参数"} extend为扩展信息作为回传参数的key
	_params string
}

// NewAlibabaAliqinTaNumberSinglecallbyvoiceRequest 初始化AlibabaAliqinTaNumberSinglecallbyvoiceAPIRequest对象
func NewAlibabaAliqinTaNumberSinglecallbyvoiceRequest() *AlibabaAliqinTaNumberSinglecallbyvoiceAPIRequest {
	return &AlibabaAliqinTaNumberSinglecallbyvoiceAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAliqinTaNumberSinglecallbyvoiceAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.ta.number.singlecallbyvoice"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAliqinTaNumberSinglecallbyvoiceAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is CalledNum Setter
// 单呼号码
func (r *AlibabaAliqinTaNumberSinglecallbyvoiceAPIRequest) SetCalledNum(_calledNum string) error {
	r._calledNum = _calledNum
	r.Set("called_num", _calledNum)
	return nil
}

// Get CalledNum Getter
func (r AlibabaAliqinTaNumberSinglecallbyvoiceAPIRequest) GetCalledNum() string {
	return r._calledNum
}

// Set is CalledShowNum Setter
// 显示号码
func (r *AlibabaAliqinTaNumberSinglecallbyvoiceAPIRequest) SetCalledShowNum(_calledShowNum string) error {
	r._calledShowNum = _calledShowNum
	r.Set("called_show_num", _calledShowNum)
	return nil
}

// Get CalledShowNum Getter
func (r AlibabaAliqinTaNumberSinglecallbyvoiceAPIRequest) GetCalledShowNum() string {
	return r._calledShowNum
}

// Set is VoiceCode Setter
// 语音文件code
func (r *AlibabaAliqinTaNumberSinglecallbyvoiceAPIRequest) SetVoiceCode(_voiceCode string) error {
	r._voiceCode = _voiceCode
	r.Set("voice_code", _voiceCode)
	return nil
}

// Get VoiceCode Getter
func (r AlibabaAliqinTaNumberSinglecallbyvoiceAPIRequest) GetVoiceCode() string {
	return r._voiceCode
}

// Set is Params Setter
// 上下文参数 示例:{"extend":"回传参数"} extend为扩展信息作为回传参数的key
func (r *AlibabaAliqinTaNumberSinglecallbyvoiceAPIRequest) SetParams(_params string) error {
	r._params = _params
	r.Set("params", _params)
	return nil
}

// Get Params Getter
func (r AlibabaAliqinTaNumberSinglecallbyvoiceAPIRequest) GetParams() string {
	return r._params
}
