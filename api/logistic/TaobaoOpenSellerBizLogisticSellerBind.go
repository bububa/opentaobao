package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// TaobaoOpenSellerBizLogisticSellerBind 店铺授权发货注册（催发货）
// taobao.open.seller.biz.logistic.seller.bind
//
// 店铺授权发货注册（催发货）
func TaobaoOpenSellerBizLogisticSellerBind(clt *core.SDKClient, req *logistic.TaobaoOpenSellerBizLogisticSellerBindAPIRequest, resp *logistic.TaobaoOpenSellerBizLogisticSellerBindAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
