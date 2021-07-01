package tmc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTmcUserCancelAPIRequest
取消用户的消息服务 API请求
taobao.tmc.user.cancel

取消用户的消息服务 */
type TaobaoTmcUserCancelAPIRequest struct {
	model.Params
	// 用户昵称
	_nick string
	// 用户所属的平台类型，tbUIC:淘宝用户; icbu: icbu用户;ae:ae用户
	_userPlatform string
}

// New
