package tmc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTmcUserGetAPIRequest
获取用户已开通消息 API请求
taobao.tmc.user.get

查询指定用户开通的消息通道和组 */
type TaobaoTmcUserGetAPIRequest struct {
	model.Params
	// 需返回的字段列表，多个字段以半角逗号分隔。可选值：TmcUser结构体中的所有字段，一定要返回topic。
	_fields string
	// 用户昵称
	_nick string
	// 用户所属的平台类型，tbUIC:淘宝用户; icbu: icbu用户;ae:ae用户
	_userPlatform string
}

// New
