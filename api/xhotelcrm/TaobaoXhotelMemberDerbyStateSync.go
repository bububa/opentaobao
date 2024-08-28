package xhotelcrm

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelcrm"
)

// TaobaoXhotelMemberDerbyStateSync 德比侧同步卡、券状态接口
// taobao.xhotel.member.derby.state.sync
//
// 德比侧同步卡、券状态接口
func TaobaoXhotelMemberDerbyStateSync(ctx context.Context, clt *core.SDKClient, req *xhotelcrm.TaobaoXhotelMemberDerbyStateSyncAPIRequest, resp *xhotelcrm.TaobaoXhotelMemberDerbyStateSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
