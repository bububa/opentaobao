package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeProjectDynamicDelete 楼盘快讯删除
// alibaba.alihouse.newhome.project.dynamic.delete
//
// 楼盘快讯删除
func AlibabaAlihouseNewhomeProjectDynamicDelete(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeProjectDynamicDeleteAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeProjectDynamicDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
