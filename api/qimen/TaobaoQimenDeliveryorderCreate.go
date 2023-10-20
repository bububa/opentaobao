package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenDeliveryorderCreate 发货单创建接口
// taobao.qimen.deliveryorder.create
//
// taobao.qimen.deliveryorder.create
func TaobaoQimenDeliveryorderCreate(clt *core.SDKClient, req *qimen.TaobaoQimenDeliveryorderCreateAPIRequest, resp *qimen.TaobaoQimenDeliveryorderCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
