package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// TaobaoLogisticsWarehouseCapacityRuleUpdate 仓产能创建/更新
// taobao.logistics.warehouse.capacity.rule.update
//
// 仓产能创建/更新
func TaobaoLogisticsWarehouseCapacityRuleUpdate(clt *core.SDKClient, req *ascp.TaobaoLogisticsWarehouseCapacityRuleUpdateAPIRequest, session string) (*ascp.TaobaoLogisticsWarehouseCapacityRuleUpdateAPIResponse, error) {
	var resp ascp.TaobaoLogisticsWarehouseCapacityRuleUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
