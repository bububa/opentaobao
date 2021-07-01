package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAliqinTaSmsNumSendAPIRequest
短信发送 API请求
alibaba.aliqin.ta.sms.num.send

短信发送 */
type AlibabaAliqinTaSmsNumSendAPIRequest struct {
	model.Params
	// 公共回传参数
	_extend string
	// 类型，normal：短信
	_smsType string
	// 短信签名
	_smsFreeSignName string
	// 短信模板变量，AckNum是变量参数
	_smsParam string
	// 接收号码
	_recNum string
	// 短信模板CODE
	_smsTemplateCode string
	// 商家自定义扩展码
	_extendCode string
	// 商家自定义扩展名,例如店铺nick
	_extendName string
}

// New
