package fenxiao

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoFenxiaoDistributorProductQuantityGet 分销商查询产品库存
// taobao.fenxiao.distributor.product.quantity.get
//
// 分销商查询产品库存
func TaobaoFenxiaoDistributorProductQuantityGet(ctx context.Context, clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoDistributorProductQuantityGetAPIRequest, resp *fenxiao.TaobaoFenxiaoDistributorProductQuantityGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
