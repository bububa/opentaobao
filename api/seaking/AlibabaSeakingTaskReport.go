package seaking

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/seaking"
)

// AlibabaSeakingTaskReport 跳转任务发布成功商品ID回传
// alibaba.seaking.task.report
//
// 跳转任务发布成功商品ID回传
func AlibabaSeakingTaskReport(ctx context.Context, clt *core.SDKClient, req *seaking.AlibabaSeakingTaskReportAPIRequest, resp *seaking.AlibabaSeakingTaskReportAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
