package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangItemmappingBatchCreate 新建商货品关联
// alibaba.dchain.aoxiang.itemmapping.batch.create
//
// 新建商货品关联
func AlibabaDchainAoxiangItemmappingBatchCreate(clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangItemmappingBatchCreateAPIRequest, session string) (*ascp.AlibabaDchainAoxiangItemmappingBatchCreateAPIResponse, error) {
	var resp ascp.AlibabaDchainAoxiangItemmappingBatchCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
