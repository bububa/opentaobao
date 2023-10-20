package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangCombinescitemBatchUpdate 更新组合货品
// alibaba.dchain.aoxiang.combinescitem.batch.update
//
// 更新组合货品
func AlibabaDchainAoxiangCombinescitemBatchUpdate(clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangCombinescitemBatchUpdateAPIRequest, resp *ascp.AlibabaDchainAoxiangCombinescitemBatchUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
