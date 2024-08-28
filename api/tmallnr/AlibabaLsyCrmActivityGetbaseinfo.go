package tmallnr

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// AlibabaLsyCrmActivityGetbaseinfo ISV查询活动
// alibaba.lsy.crm.activity.getbaseinfo
//
// ISV查询活动
func AlibabaLsyCrmActivityGetbaseinfo(ctx context.Context, clt *core.SDKClient, req *tmallnr.AlibabaLsyCrmActivityGetbaseinfoAPIRequest, resp *tmallnr.AlibabaLsyCrmActivityGetbaseinfoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
