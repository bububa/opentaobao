package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenDeliveryorderConfirm 发货单确认接口
// taobao.qimen.deliveryorder.confirm
//
// taobao.qimen.deliveryorder.confirm
func TaobaoQimenDeliveryorderConfirm(clt *core.SDKClient, req *qimen.TaobaoQimenDeliveryorderConfirmAPIRequest, resp *qimen.TaobaoQimenDeliveryorderConfirmAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
