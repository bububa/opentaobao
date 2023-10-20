package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenDeliveryorderQuery 发货单查询接口
// taobao.qimen.deliveryorder.query
//
// ERP调用奇门的发货单查询接口，查询发货单详情
func TaobaoQimenDeliveryorderQuery(clt *core.SDKClient, req *qimen.TaobaoQimenDeliveryorderQueryAPIRequest, session string) (*qimen.TaobaoQimenDeliveryorderQueryAPIResponse, error) {
	var resp qimen.TaobaoQimenDeliveryorderQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
