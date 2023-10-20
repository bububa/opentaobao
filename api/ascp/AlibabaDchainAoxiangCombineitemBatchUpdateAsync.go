package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangCombineitemBatchUpdateAsync 组合货品新建&更新
// alibaba.dchain.aoxiang.combineitem.batch.update.async
//
// 组合货品新建&amp;更新
func AlibabaDchainAoxiangCombineitemBatchUpdateAsync(clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangCombineitemBatchUpdateAsyncAPIRequest, resp *ascp.AlibabaDchainAoxiangCombineitemBatchUpdateAsyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
