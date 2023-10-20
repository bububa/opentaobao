package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Taobaologisticswarehouseoperationupdate 仓作业能力新建/更新
// taobao.logistics.warehouse.operation.update
//
// 仓作业能力新建/更新
func Taobaologisticswarehouseoperationupdate(clt *core.SDKClient, req *ascp.TaobaologisticswarehouseoperationupdateAPIRequest, session string) (*ascp.TaobaologisticswarehouseoperationupdateAPIResponse, error) {
	var resp ascp.TaobaologisticswarehouseoperationupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
