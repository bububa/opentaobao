package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscChudaTemplateSendAPIRequest
本地生活触达模板消息发送接口 API请求
alibaba.alsc.chuda.template.send

允许三方小程序通过该api发送本地生活触达消息 */
type AlibabaAlscChudaTemplateSendAPIRequest struct {
	model.Params
	// 请求参数
	_notifyRequest *TemplateNotifyRequest
}

// New
