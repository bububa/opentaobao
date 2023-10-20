package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangItemBatchUpdateAsync 货品新建/更新接口
// alibaba.dchain.aoxiang.item.batch.update.async
//
// 货品新建/更新接口
func AlibabaDchainAoxiangItemBatchUpdateAsync(clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangItemBatchUpdateAsyncAPIRequest, resp *ascp.AlibabaDchainAoxiangItemBatchUpdateAsyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
