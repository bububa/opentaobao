package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangItemBatchUpdateAsync 货品新建/更新接口
// alibaba.dchain.aoxiang.item.batch.update.async
//
// 货品新建/更新接口
func AlibabaDchainAoxiangItemBatchUpdateAsync(clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangItemBatchUpdateAsyncAPIRequest, session string) (*ascp.AlibabaDchainAoxiangItemBatchUpdateAsyncAPIResponse, error) {
	var resp ascp.AlibabaDchainAoxiangItemBatchUpdateAsyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
