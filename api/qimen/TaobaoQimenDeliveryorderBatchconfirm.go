package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

/* TaobaoQimenDeliveryorderBatchconfirm
发货单确认接口
taobao.qimen.deliveryorder.batchconfirm

taobao.qimen.deliveryorder.batchconfirm */
func TaobaoQimenDeliveryorderBatchconfirm(clt *core.SDKClient, req *qimen.TaobaoQimenDeliveryorderBatchconfirmAPIRequest, session string) (*qimen.TaobaoQimenDeliveryorderBatchconfirmAPIResponse, error) {
	var resp qimen.TaobaoQimenDeliveryorderBatchconfirmAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
