package alisports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlisportsPassportAccountGetaccountinfoAPIRequest
获取会员信息 API请求
alibaba.alisports.passport.account.getaccountinfo

获取阿里体育会员信息 */
type AlibabaAlisportsPassportAccountGetaccountinfoAPIRequest struct {
	model.Params
	// 是否获取详情0否1是 默认0
	_needDetail int64
	// 当前时间戳
	_alispTime string
	// 业务appkey
	_alispAppKey string
	// 业务加密参数
	_alispSign string
	// 查询类型：1.用户的阿里体育id, 4.用户通过登录生成的sso_token
	_type int64
	// 要查询的值
	_value string
	// 决定返回值是否包含扩展字段
	_extInfoType string
}

// New
