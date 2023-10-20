package aliqin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinFcTtsNumSinglecallAPIRequest 文本转语音通知 API请求
// alibaba.aliqin.fc.tts.num.singlecall
//
// 向指定手机号码发起单向呼叫，将文本模板内容转化为语音播放给被叫方。使用前需要在阿里大于管理中心添加去电显示号码与文本转语音模板。
type AlibabaAliqinFcTtsNumSinglecallAPIRequest struct {
	model.Params
	// 被叫号码，支持国内手机号与固话号码,格式如下057188773344,13911112222,4001112222,95500
	_calledNum string
	// 被叫号显，传入的显示号码必须是阿里大于“管理中心-号码管理”中申请或购买的号码
	_calledShowNum string
	// TTS模板ID，传入的模板必须是在阿里大于“管理中心-语音TTS模板管理”中的可用模板
	_ttsCode string
	// 公共回传参数，在“消息返回”中会透传回该参数；举例：用户可以传入自己下级的会员ID，在消息返回时，该会员ID会包含在内，用户可以根据该会员ID识别是哪位会员使用了你的应用
	_extend string
	// 文本转语音（TTS）模板变量，传参规则{"key"："value"}，key的名字须和TTS模板中的变量名一致，多个变量之间以逗号隔开，示例：{"name":"xiaoming","code":"1234"}
	_ttsParam string
}

// NewAlibabaAliqinFcTtsNumSinglecallRequest 初始化AlibabaAliqinFcTtsNumSinglecallAPIRequest对象
func NewAlibabaAliqinFcTtsNumSinglecallRequest() *AlibabaAliqinFcTtsNumSinglecallAPIRequest {
	return &AlibabaAliqinFcTtsNumSinglecallAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAliqinFcTtsNumSinglecallAPIRequest) Reset() {
	r._calledNum = ""
	r._calledShowNum = ""
	r._ttsCode = ""
	r._extend = ""
	r._ttsParam = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAliqinFcTtsNumSinglecallAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.fc.tts.num.singlecall"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAliqinFcTtsNumSinglecallAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAliqinFcTtsNumSinglecallAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCalledNum is CalledNum Setter
// 被叫号码，支持国内手机号与固话号码,格式如下057188773344,13911112222,4001112222,95500
func (r *AlibabaAliqinFcTtsNumSinglecallAPIRequest) SetCalledNum(_calledNum string) error {
	r._calledNum = _calledNum
	r.Set("called_num", _calledNum)
	return nil
}

// GetCalledNum CalledNum Getter
func (r AlibabaAliqinFcTtsNumSinglecallAPIRequest) GetCalledNum() string {
	return r._calledNum
}

// SetCalledShowNum is CalledShowNum Setter
// 被叫号显，传入的显示号码必须是阿里大于“管理中心-号码管理”中申请或购买的号码
func (r *AlibabaAliqinFcTtsNumSinglecallAPIRequest) SetCalledShowNum(_calledShowNum string) error {
	r._calledShowNum = _calledShowNum
	r.Set("called_show_num", _calledShowNum)
	return nil
}

// GetCalledShowNum CalledShowNum Getter
func (r AlibabaAliqinFcTtsNumSinglecallAPIRequest) GetCalledShowNum() string {
	return r._calledShowNum
}

// SetTtsCode is TtsCode Setter
// TTS模板ID，传入的模板必须是在阿里大于“管理中心-语音TTS模板管理”中的可用模板
func (r *AlibabaAliqinFcTtsNumSinglecallAPIRequest) SetTtsCode(_ttsCode string) error {
	r._ttsCode = _ttsCode
	r.Set("tts_code", _ttsCode)
	return nil
}

// GetTtsCode TtsCode Getter
func (r AlibabaAliqinFcTtsNumSinglecallAPIRequest) GetTtsCode() string {
	return r._ttsCode
}

// SetExtend is Extend Setter
// 公共回传参数，在“消息返回”中会透传回该参数；举例：用户可以传入自己下级的会员ID，在消息返回时，该会员ID会包含在内，用户可以根据该会员ID识别是哪位会员使用了你的应用
func (r *AlibabaAliqinFcTtsNumSinglecallAPIRequest) SetExtend(_extend string) error {
	r._extend = _extend
	r.Set("extend", _extend)
	return nil
}

// GetExtend Extend Getter
func (r AlibabaAliqinFcTtsNumSinglecallAPIRequest) GetExtend() string {
	return r._extend
}

// SetTtsParam is TtsParam Setter
// 文本转语音（TTS）模板变量，传参规则{&#34;key&#34;：&#34;value&#34;}，key的名字须和TTS模板中的变量名一致，多个变量之间以逗号隔开，示例：{&#34;name&#34;:&#34;xiaoming&#34;,&#34;code&#34;:&#34;1234&#34;}
func (r *AlibabaAliqinFcTtsNumSinglecallAPIRequest) SetTtsParam(_ttsParam string) error {
	r._ttsParam = _ttsParam
	r.Set("tts_param", _ttsParam)
	return nil
}

// GetTtsParam TtsParam Getter
func (r AlibabaAliqinFcTtsNumSinglecallAPIRequest) GetTtsParam() string {
	return r._ttsParam
}

var poolAlibabaAliqinFcTtsNumSinglecallAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAliqinFcTtsNumSinglecallRequest()
	},
}

// GetAlibabaAliqinFcTtsNumSinglecallRequest 从 sync.Pool 获取 AlibabaAliqinFcTtsNumSinglecallAPIRequest
func GetAlibabaAliqinFcTtsNumSinglecallAPIRequest() *AlibabaAliqinFcTtsNumSinglecallAPIRequest {
	return poolAlibabaAliqinFcTtsNumSinglecallAPIRequest.Get().(*AlibabaAliqinFcTtsNumSinglecallAPIRequest)
}

// ReleaseAlibabaAliqinFcTtsNumSinglecallAPIRequest 将 AlibabaAliqinFcTtsNumSinglecallAPIRequest 放入 sync.Pool
func ReleaseAlibabaAliqinFcTtsNumSinglecallAPIRequest(v *AlibabaAliqinFcTtsNumSinglecallAPIRequest) {
	v.Reset()
	poolAlibabaAliqinFcTtsNumSinglecallAPIRequest.Put(v)
}
