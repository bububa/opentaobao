package qimen

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenInventorybatchQuery 商品单仓批次库存查询接口
// taobao.qimen.inventorybatch.query
//
// ERP 通过该接口查询指定商品的单仓批次库存
func TaobaoQimenInventorybatchQuery(ctx context.Context, clt *core.SDKClient, req *qimen.TaobaoQimenInventorybatchQueryAPIRequest, resp *qimen.TaobaoQimenInventorybatchQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
