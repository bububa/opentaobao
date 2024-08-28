package omniorder

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// TaobaoQimenItemsTagQuery 打标结果查询-商品维度
// taobao.qimen.items.tag.query
//
// 调用该接口，查询某个/某批商品上的标
func TaobaoQimenItemsTagQuery(ctx context.Context, clt *core.SDKClient, req *omniorder.TaobaoQimenItemsTagQueryAPIRequest, resp *omniorder.TaobaoQimenItemsTagQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
