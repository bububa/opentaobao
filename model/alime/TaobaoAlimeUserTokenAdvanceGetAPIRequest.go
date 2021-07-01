package alime

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlimeUserTokenAdvanceGetAPIRequest
获取用户免登录令牌v2 API请求
taobao.alime.user.token.advance.get

根据第三账号信息获取用户的免登录令牌 */
type TaobaoAlimeUserTokenAdvanceGetAPIRequest struct {
	model.Params
	// 路由id, 一般为用户id，用于异地容灾
	_routing int64
	// 用户类型，0为普通用户，1为访客用户
	_type int64
	// 用户在第三方账号中的唯一id
	_foreignId string
	// 用户昵称
	_nick string
	// 小蜜分配给第三方账号的来源
	_source int64
	// 用户在小蜜账号中的唯一id
	_id int64
	// 令牌的过期时间(时间为秒)，最大为3600
	_expires int64
}

// New
