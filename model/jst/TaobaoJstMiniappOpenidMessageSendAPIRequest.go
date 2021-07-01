package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoJstMiniappOpenidMessageSendAPIRequest
单个openId用户短信发送 API请求
taobao.jst.miniapp.openid.message.send

单个openId用户短信发送 */
type TaobaoJstMiniappOpenidMessageSendAPIRequest struct {
	model.Params
	// 短信签名
	_signName string
	// 用户openId
	_openId string
	// 短信模板
	_templateCode string
	// 短信内容
	_content string
	// 短链，替换短信内容中的${url}
	_url string
	// 商家的APPKEY，如果openId是用商家的appKey生成的则需要传递
	_sellerAppKey string
	// 活动或人群code
	_crowdCode string
	// 短信拓展码
	_extendNum string
}

// New
