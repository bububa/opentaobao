package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Taobaologisticswarehousecooperationbatchconfirm 仓合作关系确认
// taobao.logistics.warehouse.cooperation.batch.confirm
//
// 仓合作关系确认
func Taobaologisticswarehousecooperationbatchconfirm(clt *core.SDKClient, req *ascp.TaobaologisticswarehousecooperationbatchconfirmAPIRequest, session string) (*ascp.TaobaologisticswarehousecooperationbatchconfirmAPIResponse, error) {
	var resp ascp.TaobaologisticswarehousecooperationbatchconfirmAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
