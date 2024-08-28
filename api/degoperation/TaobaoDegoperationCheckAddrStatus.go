package degoperation

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/degoperation"
)

// TaobaoDegoperationCheckAddrStatus 地址
// taobao.degoperation.check.addr.status
//
// 激励
func TaobaoDegoperationCheckAddrStatus(ctx context.Context, clt *core.SDKClient, req *degoperation.TaobaoDegoperationCheckAddrStatusAPIRequest, resp *degoperation.TaobaoDegoperationCheckAddrStatusAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
