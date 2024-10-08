package openmall

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openmall"
)

// TaobaoOpenmallItemSaleareaGet 查询商品可售区域
// taobao.openmall.item.salearea.get
//
// 获取openmall商品的可售区域
func TaobaoOpenmallItemSaleareaGet(ctx context.Context, clt *core.SDKClient, req *openmall.TaobaoOpenmallItemSaleareaGetAPIRequest, resp *openmall.TaobaoOpenmallItemSaleareaGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
