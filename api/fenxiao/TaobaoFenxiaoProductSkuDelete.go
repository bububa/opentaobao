package fenxiao

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoFenxiaoProductSkuDelete 产品SKU删除接口
// taobao.fenxiao.product.sku.delete
//
// 根据sku properties删除sku数据
func TaobaoFenxiaoProductSkuDelete(ctx context.Context, clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoProductSkuDeleteAPIRequest, resp *fenxiao.TaobaoFenxiaoProductSkuDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
