package cainiaohandover

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaohandover"
)

// CainiaoGlobalLogisticOrderCreate 创建物流订单
// cainiao.global.logistic.order.create
//
// 创建物流订单
func CainiaoGlobalLogisticOrderCreate(clt *core.SDKClient, req *cainiaohandover.CainiaoGlobalLogisticOrderCreateAPIRequest, resp *cainiaohandover.CainiaoGlobalLogisticOrderCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
