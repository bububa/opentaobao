package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAliqinTaNumberSinglecallbyttsAPIRequest
根据号码tts单呼 API请求
alibaba.aliqin.ta.number.singlecallbytts

将语音验证码和语音通知发布至聚石塔渠道 */
type AlibabaAliqinTaNumberSinglecallbyttsAPIRequest struct {
	model.Params
	// 被叫号码
	_calledNum string
	// 显示号码
	_calledShowNum string
	// tts文本模板code
	_ttsCode string
	// 上下文参数,tts模板含有变量的, 此参数需填写。示例:{"date":"2015年 " ,"name":"测试","extend":"回传参数"} date、name 为模板里的变量名作为key,extend为扩展信息作为回传参数的key
	_params string
}

// New
