package aliqin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinFcIvrNumCallAPIRequest ivr呼叫 API请求
// alibaba.aliqin.fc.ivr.num.call
//
// ivr呼叫
type AlibabaAliqinFcIvrNumCallAPIRequest struct {
	model.Params
	// 被叫号码，支持国内手机号与固话号码,格式如下057188773344,13911112222,4001112222,95500
	_calledNumber string
	// 被叫号码侧的号码显示，传入的显示号码可以是阿里大鱼“管理中心-号码管理”中申请通过的号码。显示号码格式如下057188773344，4001112222，95500。显示号码也可以为主叫号码。
	_calledShowNumber string
	// 可选值：tts或voice。
	_useTts string
	// 当值为tts时，menu_codet填写tts模板；当值为voice时，menu_code填写语音模板编码
	_menuCode string
	// 通话超时时长，如接通后到达120秒时，通话会因为超时自动挂断。若无需设置超时时长，可不传。
	_sessionTimeOut string
	// 公共回传参数，在消息中带回
	_extend string
	// 结束语编码，当use_tts=voice时，该字段填写语音文件编码，当use_tts=tts时，该字段填写tts模板编码
	_byeCode string
	// 当use_tts=tts时，该字段可填写tts模板变量参数
	_menuArgs string
	// 按键映射事件
	_params string
	// 人工服务号码
	_serviceNumber string
	// 播放次数
	_playTimes int64
}

// NewAlibabaAliqinFcIvrNumCallRequest 初始化AlibabaAliqinFcIvrNumCallAPIRequest对象
func NewAlibabaAliqinFcIvrNumCallRequest() *AlibabaAliqinFcIvrNumCallAPIRequest {
	return &AlibabaAliqinFcIvrNumCallAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAliqinFcIvrNumCallAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.fc.ivr.num.call"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAliqinFcIvrNumCallAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetCalledNumber is CalledNumber Setter
// 被叫号码，支持国内手机号与固话号码,格式如下057188773344,13911112222,4001112222,95500
func (r *AlibabaAliqinFcIvrNumCallAPIRequest) SetCalledNumber(_calledNumber string) error {
	r._calledNumber = _calledNumber
	r.Set("called_number", _calledNumber)
	return nil
}

// GetCalledNumber CalledNumber Getter
func (r AlibabaAliqinFcIvrNumCallAPIRequest) GetCalledNumber() string {
	return r._calledNumber
}

// SetCalledShowNumber is CalledShowNumber Setter
// 被叫号码侧的号码显示，传入的显示号码可以是阿里大鱼“管理中心-号码管理”中申请通过的号码。显示号码格式如下057188773344，4001112222，95500。显示号码也可以为主叫号码。
func (r *AlibabaAliqinFcIvrNumCallAPIRequest) SetCalledShowNumber(_calledShowNumber string) error {
	r._calledShowNumber = _calledShowNumber
	r.Set("called_show_number", _calledShowNumber)
	return nil
}

// GetCalledShowNumber CalledShowNumber Getter
func (r AlibabaAliqinFcIvrNumCallAPIRequest) GetCalledShowNumber() string {
	return r._calledShowNumber
}

// SetUseTts is UseTts Setter
// 可选值：tts或voice。
func (r *AlibabaAliqinFcIvrNumCallAPIRequest) SetUseTts(_useTts string) error {
	r._useTts = _useTts
	r.Set("use_tts", _useTts)
	return nil
}

// GetUseTts UseTts Getter
func (r AlibabaAliqinFcIvrNumCallAPIRequest) GetUseTts() string {
	return r._useTts
}

// SetMenuCode is MenuCode Setter
// 当值为tts时，menu_codet填写tts模板；当值为voice时，menu_code填写语音模板编码
func (r *AlibabaAliqinFcIvrNumCallAPIRequest) SetMenuCode(_menuCode string) error {
	r._menuCode = _menuCode
	r.Set("menu_code", _menuCode)
	return nil
}

// GetMenuCode MenuCode Getter
func (r AlibabaAliqinFcIvrNumCallAPIRequest) GetMenuCode() string {
	return r._menuCode
}

// SetSessionTimeOut is SessionTimeOut Setter
// 通话超时时长，如接通后到达120秒时，通话会因为超时自动挂断。若无需设置超时时长，可不传。
func (r *AlibabaAliqinFcIvrNumCallAPIRequest) SetSessionTimeOut(_sessionTimeOut string) error {
	r._sessionTimeOut = _sessionTimeOut
	r.Set("session_time_out", _sessionTimeOut)
	return nil
}

// GetSessionTimeOut SessionTimeOut Getter
func (r AlibabaAliqinFcIvrNumCallAPIRequest) GetSessionTimeOut() string {
	return r._sessionTimeOut
}

// SetExtend is Extend Setter
// 公共回传参数，在消息中带回
func (r *AlibabaAliqinFcIvrNumCallAPIRequest) SetExtend(_extend string) error {
	r._extend = _extend
	r.Set("extend", _extend)
	return nil
}

// GetExtend Extend Getter
func (r AlibabaAliqinFcIvrNumCallAPIRequest) GetExtend() string {
	return r._extend
}

// SetByeCode is ByeCode Setter
// 结束语编码，当use_tts=voice时，该字段填写语音文件编码，当use_tts=tts时，该字段填写tts模板编码
func (r *AlibabaAliqinFcIvrNumCallAPIRequest) SetByeCode(_byeCode string) error {
	r._byeCode = _byeCode
	r.Set("bye_code", _byeCode)
	return nil
}

// GetByeCode ByeCode Getter
func (r AlibabaAliqinFcIvrNumCallAPIRequest) GetByeCode() string {
	return r._byeCode
}

// SetMenuArgs is MenuArgs Setter
// 当use_tts=tts时，该字段可填写tts模板变量参数
func (r *AlibabaAliqinFcIvrNumCallAPIRequest) SetMenuArgs(_menuArgs string) error {
	r._menuArgs = _menuArgs
	r.Set("menu_args", _menuArgs)
	return nil
}

// GetMenuArgs MenuArgs Getter
func (r AlibabaAliqinFcIvrNumCallAPIRequest) GetMenuArgs() string {
	return r._menuArgs
}

// SetParams is Params Setter
// 按键映射事件
func (r *AlibabaAliqinFcIvrNumCallAPIRequest) SetParams(_params string) error {
	r._params = _params
	r.Set("params", _params)
	return nil
}

// GetParams Params Getter
func (r AlibabaAliqinFcIvrNumCallAPIRequest) GetParams() string {
	return r._params
}

// SetServiceNumber is ServiceNumber Setter
// 人工服务号码
func (r *AlibabaAliqinFcIvrNumCallAPIRequest) SetServiceNumber(_serviceNumber string) error {
	r._serviceNumber = _serviceNumber
	r.Set("service_number", _serviceNumber)
	return nil
}

// GetServiceNumber ServiceNumber Getter
func (r AlibabaAliqinFcIvrNumCallAPIRequest) GetServiceNumber() string {
	return r._serviceNumber
}

// SetPlayTimes is PlayTimes Setter
// 播放次数
func (r *AlibabaAliqinFcIvrNumCallAPIRequest) SetPlayTimes(_playTimes int64) error {
	r._playTimes = _playTimes
	r.Set("play_times", _playTimes)
	return nil
}

// GetPlayTimes PlayTimes Getter
func (r AlibabaAliqinFcIvrNumCallAPIRequest) GetPlayTimes() int64 {
	return r._playTimes
}
