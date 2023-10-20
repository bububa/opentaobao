package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleGoosefishUserInfoQuery 闲鱼三方容器用户信息获取
// alibaba.idle.goosefish.user.info.query
//
// 闲鱼三方容器用户信息获取
func AlibabaIdleGoosefishUserInfoQuery(clt *core.SDKClient, req *idle.AlibabaIdleGoosefishUserInfoQueryAPIRequest, resp *idle.AlibabaIdleGoosefishUserInfoQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
