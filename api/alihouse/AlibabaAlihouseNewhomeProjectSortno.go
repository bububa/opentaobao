package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeProjectSortno 新房排序值同步
// alibaba.alihouse.newhome.project.sortno
//
// 新房排序值同步
func AlibabaAlihouseNewhomeProjectSortno(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeProjectSortnoAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeProjectSortnoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
