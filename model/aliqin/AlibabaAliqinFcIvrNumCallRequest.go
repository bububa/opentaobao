package aliqin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
ivr呼叫 API请求
alibaba.aliqin.fc.ivr.num.call

ivr呼叫
*/
type AlibabaAliqinFcIvrNumCallRequest struct {
    model.Params
    // 被叫号码，支持国内手机号与固话号码,格式如下057188773344,13911112222,4001112222,95500
    _calledNumber   string
    // 被叫号码侧的号码显示，传入的显示号码可以是阿里大鱼“管理中心-号码管理”中申请通过的号码。显示号码格式如下057188773344，4001112222，95500。显示号码也可以为主叫号码。
    _calledShowNumber   string
    // 可选值：tts或voice。
    _useTts   string
    // 当值为tts时，menu_codet填写tts模板；当值为voice时，menu_code填写语音模板编码
    _menuCode   string
    // 通话超时时长，如接通后到达120秒时，通话会因为超时自动挂断。若无需设置超时时长，可不传。
    _sessionTimeOut   string
    // 公共回传参数，在消息中带回
    _extend   string
    // 结束语编码，当use_tts=voice时，该字段填写语音文件编码，当use_tts=tts时，该字段填写tts模板编码
    _byeCode   string
    // 当use_tts=tts时，该字段可填写tts模板变量参数
    _menuArgs   string
    // 播放次数
    _playTimes   int64
    // 按键映射事件
    _params   string
    // 人工服务号码
    _serviceNumber   string
}

// 初始化AlibabaAliqinFcIvrNumCallRequest对象
func NewAlibabaAliqinFcIvrNumCallRequest() *AlibabaAliqinFcIvrNumCallRequest{
    return &AlibabaAliqinFcIvrNumCallRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAliqinFcIvrNumCallRequest) GetApiMethodName() string {
    return "alibaba.aliqin.fc.ivr.num.call"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAliqinFcIvrNumCallRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CalledNumber Setter
// 被叫号码，支持国内手机号与固话号码,格式如下057188773344,13911112222,4001112222,95500
func (r *AlibabaAliqinFcIvrNumCallRequest) SetCalledNumber(_calledNumber string) error {
    r._calledNumber = _calledNumber
    r.Set("called_number", _calledNumber)
    return nil
}

// CalledNumber Getter
func (r AlibabaAliqinFcIvrNumCallRequest) GetCalledNumber() string {
    return r._calledNumber
}
// CalledShowNumber Setter
// 被叫号码侧的号码显示，传入的显示号码可以是阿里大鱼“管理中心-号码管理”中申请通过的号码。显示号码格式如下057188773344，4001112222，95500。显示号码也可以为主叫号码。
func (r *AlibabaAliqinFcIvrNumCallRequest) SetCalledShowNumber(_calledShowNumber string) error {
    r._calledShowNumber = _calledShowNumber
    r.Set("called_show_number", _calledShowNumber)
    return nil
}

// CalledShowNumber Getter
func (r AlibabaAliqinFcIvrNumCallRequest) GetCalledShowNumber() string {
    return r._calledShowNumber
}
// UseTts Setter
// 可选值：tts或voice。
func (r *AlibabaAliqinFcIvrNumCallRequest) SetUseTts(_useTts string) error {
    r._useTts = _useTts
    r.Set("use_tts", _useTts)
    return nil
}

// UseTts Getter
func (r AlibabaAliqinFcIvrNumCallRequest) GetUseTts() string {
    return r._useTts
}
// MenuCode Setter
// 当值为tts时，menu_codet填写tts模板；当值为voice时，menu_code填写语音模板编码
func (r *AlibabaAliqinFcIvrNumCallRequest) SetMenuCode(_menuCode string) error {
    r._menuCode = _menuCode
    r.Set("menu_code", _menuCode)
    return nil
}

// MenuCode Getter
func (r AlibabaAliqinFcIvrNumCallRequest) GetMenuCode() string {
    return r._menuCode
}
// SessionTimeOut Setter
// 通话超时时长，如接通后到达120秒时，通话会因为超时自动挂断。若无需设置超时时长，可不传。
func (r *AlibabaAliqinFcIvrNumCallRequest) SetSessionTimeOut(_sessionTimeOut string) error {
    r._sessionTimeOut = _sessionTimeOut
    r.Set("session_time_out", _sessionTimeOut)
    return nil
}

// SessionTimeOut Getter
func (r AlibabaAliqinFcIvrNumCallRequest) GetSessionTimeOut() string {
    return r._sessionTimeOut
}
// Extend Setter
// 公共回传参数，在消息中带回
func (r *AlibabaAliqinFcIvrNumCallRequest) SetExtend(_extend string) error {
    r._extend = _extend
    r.Set("extend", _extend)
    return nil
}

// Extend Getter
func (r AlibabaAliqinFcIvrNumCallRequest) GetExtend() string {
    return r._extend
}
// ByeCode Setter
// 结束语编码，当use_tts=voice时，该字段填写语音文件编码，当use_tts=tts时，该字段填写tts模板编码
func (r *AlibabaAliqinFcIvrNumCallRequest) SetByeCode(_byeCode string) error {
    r._byeCode = _byeCode
    r.Set("bye_code", _byeCode)
    return nil
}

// ByeCode Getter
func (r AlibabaAliqinFcIvrNumCallRequest) GetByeCode() string {
    return r._byeCode
}
// MenuArgs Setter
// 当use_tts=tts时，该字段可填写tts模板变量参数
func (r *AlibabaAliqinFcIvrNumCallRequest) SetMenuArgs(_menuArgs string) error {
    r._menuArgs = _menuArgs
    r.Set("menu_args", _menuArgs)
    return nil
}

// MenuArgs Getter
func (r AlibabaAliqinFcIvrNumCallRequest) GetMenuArgs() string {
    return r._menuArgs
}
// PlayTimes Setter
// 播放次数
func (r *AlibabaAliqinFcIvrNumCallRequest) SetPlayTimes(_playTimes int64) error {
    r._playTimes = _playTimes
    r.Set("play_times", _playTimes)
    return nil
}

// PlayTimes Getter
func (r AlibabaAliqinFcIvrNumCallRequest) GetPlayTimes() int64 {
    return r._playTimes
}
// Params Setter
// 按键映射事件
func (r *AlibabaAliqinFcIvrNumCallRequest) SetParams(_params string) error {
    r._params = _params
    r.Set("params", _params)
    return nil
}

// Params Getter
func (r AlibabaAliqinFcIvrNumCallRequest) GetParams() string {
    return r._params
}
// ServiceNumber Setter
// 人工服务号码
func (r *AlibabaAliqinFcIvrNumCallRequest) SetServiceNumber(_serviceNumber string) error {
    r._serviceNumber = _serviceNumber
    r.Set("service_number", _serviceNumber)
    return nil
}

// ServiceNumber Getter
func (r AlibabaAliqinFcIvrNumCallRequest) GetServiceNumber() string {
    return r._serviceNumber
}
