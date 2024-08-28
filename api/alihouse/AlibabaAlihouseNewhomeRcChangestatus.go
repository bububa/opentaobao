package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeRcChangestatus 图文草稿状态更新
// alibaba.alihouse.newhome.rc.changestatus
//
// 图文草稿状态更新
func AlibabaAlihouseNewhomeRcChangestatus(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeRcChangestatusAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeRcChangestatusAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
