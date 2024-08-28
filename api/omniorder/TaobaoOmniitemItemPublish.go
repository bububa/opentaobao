package omniorder

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// TaobaoOmniitemItemPublish 全渠道门店商品轻发布
// taobao.omniitem.item.publish
//
// 全渠道门店商品轻发布
func TaobaoOmniitemItemPublish(ctx context.Context, clt *core.SDKClient, req *omniorder.TaobaoOmniitemItemPublishAPIRequest, resp *omniorder.TaobaoOmniitemItemPublishAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
