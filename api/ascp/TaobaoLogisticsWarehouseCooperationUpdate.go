package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// TaobaoLogisticsWarehouseCooperationUpdate 合作商家信息同步
// taobao.logistics.warehouse.cooperation.update
//
// 合作商家信息同步
func TaobaoLogisticsWarehouseCooperationUpdate(clt *core.SDKClient, req *ascp.TaobaoLogisticsWarehouseCooperationUpdateAPIRequest, session string) (*ascp.TaobaoLogisticsWarehouseCooperationUpdateAPIResponse, error) {
	var resp ascp.TaobaoLogisticsWarehouseCooperationUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
