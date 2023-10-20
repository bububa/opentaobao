package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangCombinescitemBatchCreate 新建组合货品
// alibaba.dchain.aoxiang.combinescitem.batch.create
//
// 新建组合货品
func AlibabaDchainAoxiangCombinescitemBatchCreate(clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangCombinescitemBatchCreateAPIRequest, resp *ascp.AlibabaDchainAoxiangCombinescitemBatchCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
