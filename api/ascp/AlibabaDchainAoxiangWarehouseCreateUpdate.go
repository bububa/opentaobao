package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangWarehouseCreateUpdate 仓库信息同步
// alibaba.dchain.aoxiang.warehouse.create.update
//
// 仓库信息同步
func AlibabaDchainAoxiangWarehouseCreateUpdate(clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangWarehouseCreateUpdateAPIRequest, session string) (*ascp.AlibabaDchainAoxiangWarehouseCreateUpdateAPIResponse, error) {
	var resp ascp.AlibabaDchainAoxiangWarehouseCreateUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
