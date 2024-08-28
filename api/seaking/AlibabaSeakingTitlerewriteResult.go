package seaking

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/seaking"
)

// AlibabaSeakingTitlerewriteResult 获取标题改写任务结果
// alibaba.seaking.titlerewrite.result
//
// 获取标题改写任务结果
func AlibabaSeakingTitlerewriteResult(ctx context.Context, clt *core.SDKClient, req *seaking.AlibabaSeakingTitlerewriteResultAPIRequest, resp *seaking.AlibabaSeakingTitlerewriteResultAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
