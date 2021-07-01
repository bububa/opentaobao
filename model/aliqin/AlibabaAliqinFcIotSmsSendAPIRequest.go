package aliqin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAliqinFcIotSmsSendAPIRequest
物联网短信发送 API请求
alibaba.aliqin.fc.iot.sms.send

发送物联网短信，只允许使用物联网短信模板 */
type AlibabaAliqinFcIotSmsSendAPIRequest struct {
	model.Params
	// 公共回传参数，在“消息返回”中会透传回该参数；举例：用户可以传入自己下级的会员ID，在消息返回时，该会员ID会包含在内，用户可以根据该会员ID识别是哪位会员使用了你的应用
	_extend string
	// 短信类型，传入值请填写normal
	_smsType string
	// 短信模板变量，传参规则{"key":"value"}，key的名字须和申请模板中的变量名一致，多个变量之间以逗号隔开。示例：针对模板“验证码${code}，您正在进行${product}身份验证，打死不要告诉别人哦！”，传参时需传入{"code":"1234","product":"alidayu"}
	_smsParam string
	// 短信接收号码。
	_recNum string
	// 短信模板ID，传入的模板必须是在阿里大于“管理中心-短信模板管理”中的可用模板。示例：SMS_585014
	_smsTemplateCode string
}

// New
