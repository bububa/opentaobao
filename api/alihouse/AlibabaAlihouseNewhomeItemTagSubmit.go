package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeItemTagSubmit ETC上翻商品打标接口
// alibaba.alihouse.newhome.item.tag.submit
//
// ETC上翻商品打标接口
func AlibabaAlihouseNewhomeItemTagSubmit(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeItemTagSubmitAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeItemTagSubmitAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
