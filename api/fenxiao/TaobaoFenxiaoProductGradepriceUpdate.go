package fenxiao

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoFenxiaoProductGradepriceUpdate 根据sku设置折扣价
// taobao.fenxiao.product.gradeprice.update
//
// 供应商可以针对产品不同的sku，指定对应交易类型（代销or经销）方式下，设定折扣方式（按等级or指定分销商）以及对应优惠后的采购价格
func TaobaoFenxiaoProductGradepriceUpdate(ctx context.Context, clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoProductGradepriceUpdateAPIRequest, resp *fenxiao.TaobaoFenxiaoProductGradepriceUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
