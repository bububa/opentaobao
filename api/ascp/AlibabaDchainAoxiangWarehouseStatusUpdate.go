package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Alibabadchainaoxiangwarehousestatusupdate 启用/停用仓资源
// alibaba.dchain.aoxiang.warehouse.status.update
//
// 启用/停用仓资源
func Alibabadchainaoxiangwarehousestatusupdate(clt *core.SDKClient, req *ascp.AlibabadchainaoxiangwarehousestatusupdateAPIRequest, session string) (*ascp.AlibabadchainaoxiangwarehousestatusupdateAPIResponse, error) {
	var resp ascp.AlibabadchainaoxiangwarehousestatusupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
