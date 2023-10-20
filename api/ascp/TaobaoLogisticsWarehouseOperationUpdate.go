package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// TaobaoLogisticsWarehouseOperationUpdate 仓作业能力新建/更新
// taobao.logistics.warehouse.operation.update
//
// 仓作业能力新建/更新
func TaobaoLogisticsWarehouseOperationUpdate(clt *core.SDKClient, req *ascp.TaobaoLogisticsWarehouseOperationUpdateAPIRequest, resp *ascp.TaobaoLogisticsWarehouseOperationUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
