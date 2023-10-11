package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// TaobaoLogisticsWarehouseCooperationBatchConfirm 仓合作关系确认
// taobao.logistics.warehouse.cooperation.batch.confirm
//
// 仓合作关系确认
func TaobaoLogisticsWarehouseCooperationBatchConfirm(clt *core.SDKClient, req *ascp.TaobaoLogisticsWarehouseCooperationBatchConfirmAPIRequest, session string) (*ascp.TaobaoLogisticsWarehouseCooperationBatchConfirmAPIResponse, error) {
	var resp ascp.TaobaoLogisticsWarehouseCooperationBatchConfirmAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
