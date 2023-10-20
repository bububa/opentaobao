package ju

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ju"
)

// AlibabaJhsCommunityWechatLogin 聚划算用增淘外社群登录
// alibaba.jhs.community.wechat.login
//
// 聚划算用增淘外社群登录
func AlibabaJhsCommunityWechatLogin(clt *core.SDKClient, req *ju.AlibabaJhsCommunityWechatLoginAPIRequest, resp *ju.AlibabaJhsCommunityWechatLoginAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
