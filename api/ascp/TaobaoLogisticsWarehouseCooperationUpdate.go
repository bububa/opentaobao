package ascp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// TaobaoLogisticsWarehouseCooperationUpdate 合作商家信息同步
// taobao.logistics.warehouse.cooperation.update
//
// 合作商家信息同步
func TaobaoLogisticsWarehouseCooperationUpdate(ctx context.Context, clt *core.SDKClient, req *ascp.TaobaoLogisticsWarehouseCooperationUpdateAPIRequest, resp *ascp.TaobaoLogisticsWarehouseCooperationUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
