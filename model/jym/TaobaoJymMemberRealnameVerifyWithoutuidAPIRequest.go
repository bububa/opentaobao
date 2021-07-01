package jym

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoJymMemberRealnameVerifyWithoutuidAPIRequest
用户实名认证 API请求
taobao.jym.member.realname.verify.withoutuid

开放用户实名认证接口使用 */
type TaobaoJymMemberRealnameVerifyWithoutuidAPIRequest struct {
	model.Params
	// 加密名字串
	_encryptName string
	// 加密身份证串
	_encryptIdNo string
}

// New
