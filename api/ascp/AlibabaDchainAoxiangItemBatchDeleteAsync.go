package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangItemBatchDeleteAsync 货品与组合货品删除
// alibaba.dchain.aoxiang.item.batch.delete.async
//
// 货品与组合货品删除
func AlibabaDchainAoxiangItemBatchDeleteAsync(clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangItemBatchDeleteAsyncAPIRequest, session string) (*ascp.AlibabaDchainAoxiangItemBatchDeleteAsyncAPIResponse, error) {
	var resp ascp.AlibabaDchainAoxiangItemBatchDeleteAsyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
