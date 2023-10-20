package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangCombinescitemBatchCreate 新建组合货品
// alibaba.dchain.aoxiang.combinescitem.batch.create
//
// 新建组合货品
func AlibabaDchainAoxiangCombinescitemBatchCreate(clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangCombinescitemBatchCreateAPIRequest, session string) (*ascp.AlibabaDchainAoxiangCombinescitemBatchCreateAPIResponse, error) {
	var resp ascp.AlibabaDchainAoxiangCombinescitemBatchCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
