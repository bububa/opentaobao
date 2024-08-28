package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeProjectAdviserDelete 删除楼盘顾问
// alibaba.alihouse.newhome.project.adviser.delete
//
// 删除楼盘顾问
func AlibabaAlihouseNewhomeProjectAdviserDelete(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeProjectAdviserDeleteAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeProjectAdviserDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
