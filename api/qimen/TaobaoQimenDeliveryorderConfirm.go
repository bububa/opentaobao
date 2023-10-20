package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenDeliveryorderConfirm 发货单确认接口
// taobao.qimen.deliveryorder.confirm
//
// taobao.qimen.deliveryorder.confirm
func TaobaoQimenDeliveryorderConfirm(clt *core.SDKClient, req *qimen.TaobaoQimenDeliveryorderConfirmAPIRequest, session string) (*qimen.TaobaoQimenDeliveryorderConfirmAPIResponse, error) {
	var resp qimen.TaobaoQimenDeliveryorderConfirmAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
