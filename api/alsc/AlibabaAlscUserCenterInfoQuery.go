package alsc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscUserCenterInfoQuery 查询授权的用户信息
// alibaba.alsc.user.center.info.query
//
// 获取授权的饿了么用户信息
func AlibabaAlscUserCenterInfoQuery(ctx context.Context, clt *core.SDKClient, req *alsc.AlibabaAlscUserCenterInfoQueryAPIRequest, resp *alsc.AlibabaAlscUserCenterInfoQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
