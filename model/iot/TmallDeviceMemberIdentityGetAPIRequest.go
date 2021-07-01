package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallDeviceMemberIdentityGetAPIRequest
智能硬件会员判断 API请求
tmall.device.member.identity.get

用来识别该用户是否是商家会员· */
type TmallDeviceMemberIdentityGetAPIRequest struct {
	model.Params
	// 扩展参数为JSON字符串，用于埋点统计，source为来源字段固定值 paiyangji代表来源派样机类型设备，deviceId 为设备id，itemId 相关商品id
	_extraInfo string
	// 混淆昵称，
	_mixNick string
	// 明文nick，可不填，直接填混淆昵称
	_nick string
}

// New
