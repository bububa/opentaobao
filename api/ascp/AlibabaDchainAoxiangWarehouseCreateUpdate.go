package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Alibabadchainaoxiangwarehousecreateupdate 仓库信息同步
// alibaba.dchain.aoxiang.warehouse.create.update
//
// 仓库信息同步
func Alibabadchainaoxiangwarehousecreateupdate(clt *core.SDKClient, req *ascp.AlibabadchainaoxiangwarehousecreateupdateAPIRequest, session string) (*ascp.AlibabadchainaoxiangwarehousecreateupdateAPIResponse, error) {
	var resp ascp.AlibabadchainaoxiangwarehousecreateupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
