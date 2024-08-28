package icbulogistics

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbulogistics"
)

// AlibabaOnetouchLogisticsExpressSpecialProductTypeList 获取商品类型配置项
// alibaba.onetouch.logistics.express.special.product.type.list
//
// 获取商品类型配置项
func AlibabaOnetouchLogisticsExpressSpecialProductTypeList(ctx context.Context, clt *core.SDKClient, req *icbulogistics.AlibabaOnetouchLogisticsExpressSpecialProductTypeListAPIRequest, resp *icbulogistics.AlibabaOnetouchLogisticsExpressSpecialProductTypeListAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
