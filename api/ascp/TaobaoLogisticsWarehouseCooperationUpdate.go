package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// TaobaoLogisticsWarehouseCooperationUpdate 合作商家信息同步
// taobao.logistics.warehouse.cooperation.update
//
// 合作商家信息同步
func TaobaoLogisticsWarehouseCooperationUpdate(clt *core.SDKClient, req *ascp.TaobaoLogisticsWarehouseCooperationUpdateAPIRequest, resp *ascp.TaobaoLogisticsWarehouseCooperationUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
