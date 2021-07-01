package aliqin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAliqinFcTtsNumSinglecallAPIRequest
文本转语音通知 API请求
alibaba.aliqin.fc.tts.num.singlecall

向指定手机号码发起单向呼叫，将文本模板内容转化为语音播放给被叫方。使用前需要在阿里大于管理中心添加去电显示号码与文本转语音模板。 */
type AlibabaAliqinFcTtsNumSinglecallAPIRequest struct {
	model.Params
	// 公共回传参数，在“消息返回”中会透传回该参数；举例：用户可以传入自己下级的会员ID，在消息返回时，该会员ID会包含在内，用户可以根据该会员ID识别是哪位会员使用了你的应用
	_extend string
	// 文本转语音（TTS）模板变量，传参规则{"key"："value"}，key的名字须和TTS模板中的变量名一致，多个变量之间以逗号隔开，示例：{"name":"xiaoming","code":"1234"}
	_ttsParam string
	// 被叫号码，支持国内手机号与固话号码,格式如下057188773344,13911112222,4001112222,95500
	_calledNum string
	// 被叫号显，传入的显示号码必须是阿里大于“管理中心-号码管理”中申请或购买的号码
	_calledShowNum string
	// TTS模板ID，传入的模板必须是在阿里大于“管理中心-语音TTS模板管理”中的可用模板
	_ttsCode string
}

// New
