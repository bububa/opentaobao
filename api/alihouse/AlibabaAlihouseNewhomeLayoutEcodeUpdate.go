package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeLayoutEcodeUpdate 新房户型变更E码
// alibaba.alihouse.newhome.layout.ecode.update
//
// 新房户型变更E码
func AlibabaAlihouseNewhomeLayoutEcodeUpdate(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeLayoutEcodeUpdateAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeLayoutEcodeUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
