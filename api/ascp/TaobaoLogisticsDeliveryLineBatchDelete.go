package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// TaobaoLogisticsDeliveryLineBatchDelete 线路能力删除
// taobao.logistics.delivery.line.batch.delete
//
// 线路能力删除
func TaobaoLogisticsDeliveryLineBatchDelete(clt *core.SDKClient, req *ascp.TaobaoLogisticsDeliveryLineBatchDeleteAPIRequest, session string) (*ascp.TaobaoLogisticsDeliveryLineBatchDeleteAPIResponse, error) {
	var resp ascp.TaobaoLogisticsDeliveryLineBatchDeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
