package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeProjectEcode 楼盘e码更新
// alibaba.alihouse.newhome.project.ecode
//
// 更新楼盘ecode
func AlibabaAlihouseNewhomeProjectEcode(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeProjectEcodeAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeProjectEcodeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
