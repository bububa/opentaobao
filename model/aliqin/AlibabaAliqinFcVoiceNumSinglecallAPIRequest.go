package aliqin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
语音通知 API请求
alibaba.aliqin.fc.voice.num.singlecall

向指定手机号码发起单向呼叫，播放指定的语音文件内容。使用前需要在阿里大于管理中心添加去电显示号码与语音文件。
*/
type AlibabaAliqinFcVoiceNumSinglecallAPIRequest struct {
    model.Params
    // 公共回传参数，在“消息返回”中会透传回该参数；举例：用户可以传入自己下级的会员ID，在消息返回时，该会员ID会包含在内，用户可以根据该会员ID识别是哪位会员使用了你的应用
    _extend   string
    // 被叫号码，支持国内手机号与固话号码,格式如下057188773344,13911112222,4001112222,95500
    _calledNum   string
    // 被叫号显，传入的显示号码必须是阿里大于“管理中心-号码管理”中申请通过的号码
    _calledShowNum   string
    // 语音文件ID，传入的语音文件必须是在阿里大于“管理中心-语音文件管理”中的可用语音文件
    _voiceCode   string
}

// 初始化AlibabaAliqinFcVoiceNumSinglecallAPIRequest对象
func NewAlibabaAliqinFcVoiceNumSinglecallRequest() *AlibabaAliqinFcVoiceNumSinglecallAPIRequest{
    return &AlibabaAliqinFcVoiceNumSinglecallAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAliqinFcVoiceNumSinglecallAPIRequest) GetApiMethodName() string {
    return "alibaba.aliqin.fc.voice.num.singlecall"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAliqinFcVoiceNumSinglecallAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Extend Setter
// 公共回传参数，在“消息返回”中会透传回该参数；举例：用户可以传入自己下级的会员ID，在消息返回时，该会员ID会包含在内，用户可以根据该会员ID识别是哪位会员使用了你的应用
func (r *AlibabaAliqinFcVoiceNumSinglecallAPIRequest) SetExtend(_extend string) error {
    r._extend = _extend
    r.Set("extend", _extend)
    return nil
}

// Extend Getter
func (r AlibabaAliqinFcVoiceNumSinglecallAPIRequest) GetExtend() string {
    return r._extend
}
// CalledNum Setter
// 被叫号码，支持国内手机号与固话号码,格式如下057188773344,13911112222,4001112222,95500
func (r *AlibabaAliqinFcVoiceNumSinglecallAPIRequest) SetCalledNum(_calledNum string) error {
    r._calledNum = _calledNum
    r.Set("called_num", _calledNum)
    return nil
}

// CalledNum Getter
func (r AlibabaAliqinFcVoiceNumSinglecallAPIRequest) GetCalledNum() string {
    return r._calledNum
}
// CalledShowNum Setter
// 被叫号显，传入的显示号码必须是阿里大于“管理中心-号码管理”中申请通过的号码
func (r *AlibabaAliqinFcVoiceNumSinglecallAPIRequest) SetCalledShowNum(_calledShowNum string) error {
    r._calledShowNum = _calledShowNum
    r.Set("called_show_num", _calledShowNum)
    return nil
}

// CalledShowNum Getter
func (r AlibabaAliqinFcVoiceNumSinglecallAPIRequest) GetCalledShowNum() string {
    return r._calledShowNum
}
// VoiceCode Setter
// 语音文件ID，传入的语音文件必须是在阿里大于“管理中心-语音文件管理”中的可用语音文件
func (r *AlibabaAliqinFcVoiceNumSinglecallAPIRequest) SetVoiceCode(_voiceCode string) error {
    r._voiceCode = _voiceCode
    r.Set("voice_code", _voiceCode)
    return nil
}

// VoiceCode Getter
func (r AlibabaAliqinFcVoiceNumSinglecallAPIRequest) GetVoiceCode() string {
    return r._voiceCode
}
