package logistic

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// TaobaoOpenSellerBizLogisticSellerBind 店铺授权发货注册（催发货）
// taobao.open.seller.biz.logistic.seller.bind
//
// 店铺授权发货注册（催发货）
func TaobaoOpenSellerBizLogisticSellerBind(ctx context.Context, clt *core.SDKClient, req *logistic.TaobaoOpenSellerBizLogisticSellerBindAPIRequest, resp *logistic.TaobaoOpenSellerBizLogisticSellerBindAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
