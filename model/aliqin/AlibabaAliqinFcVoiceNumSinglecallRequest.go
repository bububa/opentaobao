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
type AlibabaAliqinFcVoiceNumSinglecallRequest struct {
    model.Params
    // 公共回传参数，在“消息返回”中会透传回该参数；举例：用户可以传入自己下级的会员ID，在消息返回时，该会员ID会包含在内，用户可以根据该会员ID识别是哪位会员使用了你的应用
    extend   string
    // 被叫号码，支持国内手机号与固话号码,格式如下057188773344,13911112222,4001112222,95500
    calledNum   string
    // 被叫号显，传入的显示号码必须是阿里大于“管理中心-号码管理”中申请通过的号码
    calledShowNum   string
    // 语音文件ID，传入的语音文件必须是在阿里大于“管理中心-语音文件管理”中的可用语音文件
    voiceCode   string
}

// 初始化AlibabaAliqinFcVoiceNumSinglecallRequest对象
func NewAlibabaAliqinFcVoiceNumSinglecallRequest() *AlibabaAliqinFcVoiceNumSinglecallRequest{
    return &AlibabaAliqinFcVoiceNumSinglecallRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAliqinFcVoiceNumSinglecallRequest) GetApiMethodName() string {
    return "alibaba.aliqin.fc.voice.num.singlecall"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAliqinFcVoiceNumSinglecallRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Extend Setter
// 公共回传参数，在“消息返回”中会透传回该参数；举例：用户可以传入自己下级的会员ID，在消息返回时，该会员ID会包含在内，用户可以根据该会员ID识别是哪位会员使用了你的应用
func (r *AlibabaAliqinFcVoiceNumSinglecallRequest) SetExtend(extend string) error {
    r.extend = extend
    r.Set("extend", extend)
    return nil
}

// Extend Getter
func (r AlibabaAliqinFcVoiceNumSinglecallRequest) GetExtend() string {
    return r.extend
}
// CalledNum Setter
// 被叫号码，支持国内手机号与固话号码,格式如下057188773344,13911112222,4001112222,95500
func (r *AlibabaAliqinFcVoiceNumSinglecallRequest) SetCalledNum(calledNum string) error {
    r.calledNum = calledNum
    r.Set("called_num", calledNum)
    return nil
}

// CalledNum Getter
func (r AlibabaAliqinFcVoiceNumSinglecallRequest) GetCalledNum() string {
    return r.calledNum
}
// CalledShowNum Setter
// 被叫号显，传入的显示号码必须是阿里大于“管理中心-号码管理”中申请通过的号码
func (r *AlibabaAliqinFcVoiceNumSinglecallRequest) SetCalledShowNum(calledShowNum string) error {
    r.calledShowNum = calledShowNum
    r.Set("called_show_num", calledShowNum)
    return nil
}

// CalledShowNum Getter
func (r AlibabaAliqinFcVoiceNumSinglecallRequest) GetCalledShowNum() string {
    return r.calledShowNum
}
// VoiceCode Setter
// 语音文件ID，传入的语音文件必须是在阿里大于“管理中心-语音文件管理”中的可用语音文件
func (r *AlibabaAliqinFcVoiceNumSinglecallRequest) SetVoiceCode(voiceCode string) error {
    r.voiceCode = voiceCode
    r.Set("voice_code", voiceCode)
    return nil
}

// VoiceCode Getter
func (r AlibabaAliqinFcVoiceNumSinglecallRequest) GetVoiceCode() string {
    return r.voiceCode
}
