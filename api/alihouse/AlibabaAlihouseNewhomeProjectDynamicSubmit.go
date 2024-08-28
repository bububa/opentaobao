package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeProjectDynamicSubmit 提交楼盘快讯
// alibaba.alihouse.newhome.project.dynamic.submit
//
// 提交楼盘快讯
func AlibabaAlihouseNewhomeProjectDynamicSubmit(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeProjectDynamicSubmitAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeProjectDynamicSubmitAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
