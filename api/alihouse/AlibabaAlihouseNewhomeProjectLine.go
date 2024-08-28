package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeProjectLine 楼盘上下架
// alibaba.alihouse.newhome.project.line
//
// 上下架楼盘
func AlibabaAlihouseNewhomeProjectLine(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeProjectLineAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeProjectLineAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
