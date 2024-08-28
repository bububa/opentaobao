package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeEcodeUpdate 新房货变更E码
// alibaba.alihouse.newhome.ecode.update
//
// 新房货变更E码
func AlibabaAlihouseNewhomeEcodeUpdate(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeEcodeUpdateAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeEcodeUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
