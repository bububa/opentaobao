package dutyfree

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/dutyfree"
)

// AlibabaDutyfreeStockQuery 对外库存查询接口
// alibaba.dutyfree.stock.query
//
// 对外部服务提供库存查询接口
func AlibabaDutyfreeStockQuery(ctx context.Context, clt *core.SDKClient, req *dutyfree.AlibabaDutyfreeStockQueryAPIRequest, resp *dutyfree.AlibabaDutyfreeStockQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
