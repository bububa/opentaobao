package tmallnr

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// AlibabaLsyCrmActivityGet 私域导购查询活动详情
// alibaba.lsy.crm.activity.get
//
// 私域导购查询活动详情
func AlibabaLsyCrmActivityGet(ctx context.Context, clt *core.SDKClient, req *tmallnr.AlibabaLsyCrmActivityGetAPIRequest, resp *tmallnr.AlibabaLsyCrmActivityGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
