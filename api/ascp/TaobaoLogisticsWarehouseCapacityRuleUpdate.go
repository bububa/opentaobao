package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Taobaologisticswarehousecapacityruleupdate 仓产能创建/更新
// taobao.logistics.warehouse.capacity.rule.update
//
// 仓产能创建/更新
func Taobaologisticswarehousecapacityruleupdate(clt *core.SDKClient, req *ascp.TaobaologisticswarehousecapacityruleupdateAPIRequest, session string) (*ascp.TaobaologisticswarehousecapacityruleupdateAPIResponse, error) {
	var resp ascp.TaobaologisticswarehousecapacityruleupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
