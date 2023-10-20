package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenDeliveryorderCreate 发货单创建接口
// taobao.qimen.deliveryorder.create
//
// taobao.qimen.deliveryorder.create
func TaobaoQimenDeliveryorderCreate(clt *core.SDKClient, req *qimen.TaobaoQimenDeliveryorderCreateAPIRequest, session string) (*qimen.TaobaoQimenDeliveryorderCreateAPIResponse, error) {
	var resp qimen.TaobaoQimenDeliveryorderCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
