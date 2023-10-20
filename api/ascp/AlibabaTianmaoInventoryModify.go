package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaTianmaoInventoryModify 阿里巴巴.天猫.aic库存.修改
// alibaba.tianmao.inventory.modify
//
// 阿里巴巴.天猫.aic库存.修改
func AlibabaTianmaoInventoryModify(clt *core.SDKClient, req *ascp.AlibabaTianmaoInventoryModifyAPIRequest, session string) (*ascp.AlibabaTianmaoInventoryModifyAPIResponse, error) {
	var resp ascp.AlibabaTianmaoInventoryModifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
