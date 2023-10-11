package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// TaobaoLogisticsDeliveryLineBatchUpdate 线路能力创建/更新
// taobao.logistics.delivery.line.batch.update
//
// 线路能力创建/更新
func TaobaoLogisticsDeliveryLineBatchUpdate(clt *core.SDKClient, req *ascp.TaobaoLogisticsDeliveryLineBatchUpdateAPIRequest, session string) (*ascp.TaobaoLogisticsDeliveryLineBatchUpdateAPIResponse, error) {
	var resp ascp.TaobaoLogisticsDeliveryLineBatchUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
