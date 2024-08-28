package subuser

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/subuser"
)

// TaobaoSubusersInfoQuery 根据当前子账号登陆态，获取该子账号基本信息
// taobao.subusers.info.query
//
// 根据当前子账号登陆态，获取该子账号基本信息
func TaobaoSubusersInfoQuery(ctx context.Context, clt *core.SDKClient, req *subuser.TaobaoSubusersInfoQueryAPIRequest, resp *subuser.TaobaoSubusersInfoQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
