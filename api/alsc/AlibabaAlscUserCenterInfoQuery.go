package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscUserCenterInfoQuery 查询授权的用户信息
// alibaba.alsc.user.center.info.query
//
// 获取授权的饿了么用户信息
func AlibabaAlscUserCenterInfoQuery(clt *core.SDKClient, req *alsc.AlibabaAlscUserCenterInfoQueryAPIRequest, resp *alsc.AlibabaAlscUserCenterInfoQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
