package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// TaobaoLogisticsWarehouseCooperationQuery 仓合作关系查询
// taobao.logistics.warehouse.cooperation.query
//
// 仓合作关系查询
func TaobaoLogisticsWarehouseCooperationQuery(clt *core.SDKClient, req *ascp.TaobaoLogisticsWarehouseCooperationQueryAPIRequest, session string) (*ascp.TaobaoLogisticsWarehouseCooperationQueryAPIResponse, error) {
	var resp ascp.TaobaoLogisticsWarehouseCooperationQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
