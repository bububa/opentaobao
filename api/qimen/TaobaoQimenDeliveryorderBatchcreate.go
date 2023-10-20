package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenDeliveryorderBatchcreate 发货单创建批量接口
// taobao.qimen.deliveryorder.batchcreate
//
// taobao.qimen.deliveryorder.batchcreate
func TaobaoQimenDeliveryorderBatchcreate(clt *core.SDKClient, req *qimen.TaobaoQimenDeliveryorderBatchcreateAPIRequest, session string) (*qimen.TaobaoQimenDeliveryorderBatchcreateAPIResponse, error) {
	var resp qimen.TaobaoQimenDeliveryorderBatchcreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
