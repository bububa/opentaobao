package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangItemmappingBatchCreate 新建商货品关联
// alibaba.dchain.aoxiang.itemmapping.batch.create
//
// 新建商货品关联
func AlibabaDchainAoxiangItemmappingBatchCreate(clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangItemmappingBatchCreateAPIRequest, resp *ascp.AlibabaDchainAoxiangItemmappingBatchCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
