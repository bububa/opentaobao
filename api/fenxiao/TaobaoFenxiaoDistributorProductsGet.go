package fenxiao

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoFenxiaoDistributorProductsGet 分销商查询产品信息
// taobao.fenxiao.distributor.products.get
//
// 分销商查询供应商产品信息
func TaobaoFenxiaoDistributorProductsGet(ctx context.Context, clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoDistributorProductsGetAPIRequest, resp *fenxiao.TaobaoFenxiaoDistributorProductsGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
