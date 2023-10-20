package alisports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alisports"
)

// AlibabaAlisportsPassportOauthTokenvalidate 阿里体育会员系统--获取登录信息接口
// alibaba.alisports.passport.oauth.tokenvalidate
//
// 阿里体育会员系统--获取登录信息接口
func AlibabaAlisportsPassportOauthTokenvalidate(clt *core.SDKClient, req *alisports.AlibabaAlisportsPassportOauthTokenvalidateAPIRequest, resp *alisports.AlibabaAlisportsPassportOauthTokenvalidateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
