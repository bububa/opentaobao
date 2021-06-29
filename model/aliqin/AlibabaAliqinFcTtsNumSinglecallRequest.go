package aliqin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
文本转语音通知 API请求
alibaba.aliqin.fc.tts.num.singlecall

向指定手机号码发起单向呼叫，将文本模板内容转化为语音播放给被叫方。使用前需要在阿里大于管理中心添加去电显示号码与文本转语音模板。
*/
type AlibabaAliqinFcTtsNumSinglecallRequest struct {
    model.Params
    // 公共回传参数，在“消息返回”中会透传回该参数；举例：用户可以传入自己下级的会员ID，在消息返回时，该会员ID会包含在内，用户可以根据该会员ID识别是哪位会员使用了你的应用
    extend   string
    // 文本转语音（TTS）模板变量，传参规则{"key"："value"}，key的名字须和TTS模板中的变量名一致，多个变量之间以逗号隔开，示例：{"name":"xiaoming","code":"1234"}
    ttsParam   string
    // 被叫号码，支持国内手机号与固话号码,格式如下057188773344,13911112222,4001112222,95500
    calledNum   string
    // 被叫号显，传入的显示号码必须是阿里大于“管理中心-号码管理”中申请或购买的号码
    calledShowNum   string
    // TTS模板ID，传入的模板必须是在阿里大于“管理中心-语音TTS模板管理”中的可用模板
    ttsCode   string
}

// 初始化AlibabaAliqinFcTtsNumSinglecallRequest对象
func NewAlibabaAliqinFcTtsNumSinglecallRequest() *AlibabaAliqinFcTtsNumSinglecallRequest{
    return &AlibabaAliqinFcTtsNumSinglecallRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAliqinFcTtsNumSinglecallRequest) GetApiMethodName() string {
    return "alibaba.aliqin.fc.tts.num.singlecall"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAliqinFcTtsNumSinglecallRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Extend Setter
// 公共回传参数，在“消息返回”中会透传回该参数；举例：用户可以传入自己下级的会员ID，在消息返回时，该会员ID会包含在内，用户可以根据该会员ID识别是哪位会员使用了你的应用
func (r *AlibabaAliqinFcTtsNumSinglecallRequest) SetExtend(extend string) error {
    r.extend = extend
    r.Set("extend", extend)
    return nil
}

// Extend Getter
func (r AlibabaAliqinFcTtsNumSinglecallRequest) GetExtend() string {
    return r.extend
}
// TtsParam Setter
// 文本转语音（TTS）模板变量，传参规则{"key"："value"}，key的名字须和TTS模板中的变量名一致，多个变量之间以逗号隔开，示例：{"name":"xiaoming","code":"1234"}
func (r *AlibabaAliqinFcTtsNumSinglecallRequest) SetTtsParam(ttsParam string) error {
    r.ttsParam = ttsParam
    r.Set("tts_param", ttsParam)
    return nil
}

// TtsParam Getter
func (r AlibabaAliqinFcTtsNumSinglecallRequest) GetTtsParam() string {
    return r.ttsParam
}
// CalledNum Setter
// 被叫号码，支持国内手机号与固话号码,格式如下057188773344,13911112222,4001112222,95500
func (r *AlibabaAliqinFcTtsNumSinglecallRequest) SetCalledNum(calledNum string) error {
    r.calledNum = calledNum
    r.Set("called_num", calledNum)
    return nil
}

// CalledNum Getter
func (r AlibabaAliqinFcTtsNumSinglecallRequest) GetCalledNum() string {
    return r.calledNum
}
// CalledShowNum Setter
// 被叫号显，传入的显示号码必须是阿里大于“管理中心-号码管理”中申请或购买的号码
func (r *AlibabaAliqinFcTtsNumSinglecallRequest) SetCalledShowNum(calledShowNum string) error {
    r.calledShowNum = calledShowNum
    r.Set("called_show_num", calledShowNum)
    return nil
}

// CalledShowNum Getter
func (r AlibabaAliqinFcTtsNumSinglecallRequest) GetCalledShowNum() string {
    return r.calledShowNum
}
// TtsCode Setter
// TTS模板ID，传入的模板必须是在阿里大于“管理中心-语音TTS模板管理”中的可用模板
func (r *AlibabaAliqinFcTtsNumSinglecallRequest) SetTtsCode(ttsCode string) error {
    r.ttsCode = ttsCode
    r.Set("tts_code", ttsCode)
    return nil
}

// TtsCode Getter
func (r AlibabaAliqinFcTtsNumSinglecallRequest) GetTtsCode() string {
    return r.ttsCode
}
