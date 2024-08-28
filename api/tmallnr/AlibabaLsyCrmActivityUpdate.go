package tmallnr

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// AlibabaLsyCrmActivityUpdate ISV活动修改
// alibaba.lsy.crm.activity.update
//
// ISV活动修改
func AlibabaLsyCrmActivityUpdate(ctx context.Context, clt *core.SDKClient, req *tmallnr.AlibabaLsyCrmActivityUpdateAPIRequest, resp *tmallnr.AlibabaLsyCrmActivityUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
