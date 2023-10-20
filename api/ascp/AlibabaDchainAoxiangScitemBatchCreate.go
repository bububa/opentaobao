package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangScitemBatchCreate 新建货品
// alibaba.dchain.aoxiang.scitem.batch.create
//
// 新建货品
func AlibabaDchainAoxiangScitemBatchCreate(clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangScitemBatchCreateAPIRequest, session string) (*ascp.AlibabaDchainAoxiangScitemBatchCreateAPIResponse, error) {
	var resp ascp.AlibabaDchainAoxiangScitemBatchCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
