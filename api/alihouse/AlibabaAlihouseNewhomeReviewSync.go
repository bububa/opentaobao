package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeReviewSync 天猫好房楼盘评测同步
// alibaba.alihouse.newhome.review.sync
//
// 接受楼盘测评图文信息
func AlibabaAlihouseNewhomeReviewSync(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeReviewSyncAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeReviewSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
