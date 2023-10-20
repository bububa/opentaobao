package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangItemmappingDelete 删除商货品关联关系
// alibaba.dchain.aoxiang.itemmapping.delete
//
// 删除商货品关联关系
func AlibabaDchainAoxiangItemmappingDelete(clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangItemmappingDeleteAPIRequest, resp *ascp.AlibabaDchainAoxiangItemmappingDeleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
