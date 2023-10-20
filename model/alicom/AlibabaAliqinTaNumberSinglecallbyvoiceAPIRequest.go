package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaaliqintanumbersinglecallbyvoiceAPIRequest 根据号码tts单呼 API请求
// alibaba.aliqin.ta.number.singlecallbyvoice
//
// 根据号码语音单呼
type AlibabaaliqintanumbersinglecallbyvoiceAPIRequest struct {
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

// NewAlibabaaliqintanumbersinglecallbyvoiceRequest 初始化AlibabaaliqintanumbersinglecallbyvoiceAPIRequest对象
func NewAlibabaaliqintanumbersinglecallbyvoiceRequest() *AlibabaaliqintanumbersinglecallbyvoiceAPIRequest {
	return &AlibabaaliqintanumbersinglecallbyvoiceAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaaliqintanumbersinglecallbyvoiceAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.ta.number.singlecallbyvoice"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaaliqintanumbersinglecallbyvoiceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaaliqintanumbersinglecallbyvoiceAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCalledNum is CalledNum Setter
// 单呼号码
func (r *AlibabaaliqintanumbersinglecallbyvoiceAPIRequest) SetCalledNum(_calledNum string) error {
	r._calledNum = _calledNum
	r.Set("called_num", _calledNum)
	return nil
}

// GetCalledNum CalledNum Getter
func (r AlibabaaliqintanumbersinglecallbyvoiceAPIRequest) GetCalledNum() string {
	return r._calledNum
}

// SetCalledShowNum is CalledShowNum Setter
// 显示号码
func (r *AlibabaaliqintanumbersinglecallbyvoiceAPIRequest) SetCalledShowNum(_calledShowNum string) error {
	r._calledShowNum = _calledShowNum
	r.Set("called_show_num", _calledShowNum)
	return nil
}

// GetCalledShowNum CalledShowNum Getter
func (r AlibabaaliqintanumbersinglecallbyvoiceAPIRequest) GetCalledShowNum() string {
	return r._calledShowNum
}

// SetVoiceCode is VoiceCode Setter
// 语音文件code
func (r *AlibabaaliqintanumbersinglecallbyvoiceAPIRequest) SetVoiceCode(_voiceCode string) error {
	r._voiceCode = _voiceCode
	r.Set("voice_code", _voiceCode)
	return nil
}

// GetVoiceCode VoiceCode Getter
func (r AlibabaaliqintanumbersinglecallbyvoiceAPIRequest) GetVoiceCode() string {
	return r._voiceCode
}

// SetParams is Params Setter
// 上下文参数 示例:{&#34;extend&#34;:&#34;回传参数&#34;} extend为扩展信息作为回传参数的key
func (r *AlibabaaliqintanumbersinglecallbyvoiceAPIRequest) SetParams(_params string) error {
	r._params = _params
	r.Set("params", _params)
	return nil
}

// GetParams Params Getter
func (r AlibabaaliqintanumbersinglecallbyvoiceAPIRequest) GetParams() string {
	return r._params
}
