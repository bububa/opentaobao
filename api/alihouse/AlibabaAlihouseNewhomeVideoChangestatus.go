package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeVideoChangestatus 视频草稿状态更新
// alibaba.alihouse.newhome.video.changestatus
//
// 视频草稿状态更新
func AlibabaAlihouseNewhomeVideoChangestatus(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeVideoChangestatusAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeVideoChangestatusAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
