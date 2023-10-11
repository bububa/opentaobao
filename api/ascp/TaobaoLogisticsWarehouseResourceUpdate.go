package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// TaobaoLogisticsWarehouseResourceUpdate 服务商新建/更新仓
// taobao.logistics.warehouse.resource.update
//
// 服务商新建/更新仓
func TaobaoLogisticsWarehouseResourceUpdate(clt *core.SDKClient, req *ascp.TaobaoLogisticsWarehouseResourceUpdateAPIRequest, session string) (*ascp.TaobaoLogisticsWarehouseResourceUpdateAPIResponse, error) {
	var resp ascp.TaobaoLogisticsWarehouseResourceUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
