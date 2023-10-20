package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Taobaologisticswarehousecooperationquery 仓合作关系查询
// taobao.logistics.warehouse.cooperation.query
//
// 仓合作关系查询
func Taobaologisticswarehousecooperationquery(clt *core.SDKClient, req *ascp.TaobaologisticswarehousecooperationqueryAPIRequest, session string) (*ascp.TaobaologisticswarehousecooperationqueryAPIResponse, error) {
	var resp ascp.TaobaologisticswarehousecooperationqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
