package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeSystemSeller 商品发布授权
// alibaba.alihouse.newhome.system.seller
//
// 商品发布授权
func AlibabaAlihouseNewhomeSystemSeller(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeSystemSellerAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeSystemSellerAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
