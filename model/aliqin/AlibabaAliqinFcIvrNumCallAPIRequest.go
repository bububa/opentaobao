package aliqin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAliqinFcIvrNumCallAPIRequest
ivr呼叫 API请求
alibaba.aliqin.fc.ivr.num.call

ivr呼叫 */
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
	// 播放次数
	_playTimes int64
	// 按键映射事件
	_params string
	// 人工服务号码
	_serviceNumber string
}

// New
