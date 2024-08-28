package tmallnr

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// AlibabaLsyCrmActivityPageUpdate ISV活动页面创建修改
// alibaba.lsy.crm.activity.page.update
//
// ISV活动页面创建修改
func AlibabaLsyCrmActivityPageUpdate(ctx context.Context, clt *core.SDKClient, req *tmallnr.AlibabaLsyCrmActivityPageUpdateAPIRequest, resp *tmallnr.AlibabaLsyCrmActivityPageUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
