package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangWarehouseStatusUpdate 启用/停用仓资源
// alibaba.dchain.aoxiang.warehouse.status.update
//
// 启用/停用仓资源
func AlibabaDchainAoxiangWarehouseStatusUpdate(clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangWarehouseStatusUpdateAPIRequest, session string) (*ascp.AlibabaDchainAoxiangWarehouseStatusUpdateAPIResponse, error) {
	var resp ascp.AlibabaDchainAoxiangWarehouseStatusUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
