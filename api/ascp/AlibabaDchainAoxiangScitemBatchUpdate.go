package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangScitemBatchUpdate alibaba.dchain.aoxiang.scitem.batch.update
// alibaba.dchain.aoxiang.scitem.batch.update
//
// 更新货品
func AlibabaDchainAoxiangScitemBatchUpdate(clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangScitemBatchUpdateAPIRequest, session string) (*ascp.AlibabaDchainAoxiangScitemBatchUpdateAPIResponse, error) {
	var resp ascp.AlibabaDchainAoxiangScitemBatchUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
