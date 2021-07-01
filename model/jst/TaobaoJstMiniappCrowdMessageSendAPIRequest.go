package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoJstMiniappCrowdMessageSendAPIRequest
小程序活动短信发送 API请求
taobao.jst.miniapp.crowd.message.send

小程序活动短信发送 */
type TaobaoJstMiniappCrowdMessageSendAPIRequest struct {
	model.Params
	// 短信签名
	_signName string
	// 活动code
	_crowdCode string
	// 短信模板，必须为全变量模板
	_templateCode string
	// 短信内容
	_content string
	// 短信中携带的短链，会替换短信内容中的${url}
	_url string
}

// New
