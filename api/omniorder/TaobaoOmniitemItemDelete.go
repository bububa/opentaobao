package omniorder

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// TaobaoOmniitemItemDelete 全渠道商品删除
// taobao.omniitem.item.delete
//
// 全渠道商品删除，能够对门店商品库商品进行删除动作
func TaobaoOmniitemItemDelete(ctx context.Context, clt *core.SDKClient, req *omniorder.TaobaoOmniitemItemDeleteAPIRequest, resp *omniorder.TaobaoOmniitemItemDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
