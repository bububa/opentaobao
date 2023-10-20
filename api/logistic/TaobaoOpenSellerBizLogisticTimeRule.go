package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// TaobaoOpenSellerBizLogisticTimeRule 商家自定义发货时效
// taobao.open.seller.biz.logistic.time.rule
//
// 服务商回传商家自定义发货时效
func TaobaoOpenSellerBizLogisticTimeRule(clt *core.SDKClient, req *logistic.TaobaoOpenSellerBizLogisticTimeRuleAPIRequest, resp *logistic.TaobaoOpenSellerBizLogisticTimeRuleAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
