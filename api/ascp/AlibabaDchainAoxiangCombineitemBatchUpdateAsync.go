package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangCombineitemBatchUpdateAsync 组合货品新建&更新
// alibaba.dchain.aoxiang.combineitem.batch.update.async
//
// 组合货品新建&更新
func AlibabaDchainAoxiangCombineitemBatchUpdateAsync(clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangCombineitemBatchUpdateAsyncAPIRequest, session string) (*ascp.AlibabaDchainAoxiangCombineitemBatchUpdateAsyncAPIResponse, error) {
	var resp ascp.AlibabaDchainAoxiangCombineitemBatchUpdateAsyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
