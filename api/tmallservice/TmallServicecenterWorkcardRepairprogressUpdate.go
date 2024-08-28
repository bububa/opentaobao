package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterWorkcardRepairprogressUpdate 更新维修进度
// tmall.servicecenter.workcard.repairprogress.update
//
// 提供给外部合作服务商的维修进度更改接口
func TmallServicecenterWorkcardRepairprogressUpdate(ctx context.Context, clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkcardRepairprogressUpdateAPIRequest, resp *tmallservice.TmallServicecenterWorkcardRepairprogressUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
