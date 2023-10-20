package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Taobaologisticswarehouseresourceupdate 服务商新建/更新仓
// taobao.logistics.warehouse.resource.update
//
// 服务商新建/更新仓
func Taobaologisticswarehouseresourceupdate(clt *core.SDKClient, req *ascp.TaobaologisticswarehouseresourceupdateAPIRequest, session string) (*ascp.TaobaologisticswarehouseresourceupdateAPIResponse, error) {
	var resp ascp.TaobaologisticswarehouseresourceupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
