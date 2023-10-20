package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangItemmappingUnbundle 商货关联解绑
// alibaba.dchain.aoxiang.itemmapping.unbundle
//
// 商货关联解绑
func AlibabaDchainAoxiangItemmappingUnbundle(clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangItemmappingUnbundleAPIRequest, session string) (*ascp.AlibabaDchainAoxiangItemmappingUnbundleAPIResponse, error) {
	var resp ascp.AlibabaDchainAoxiangItemmappingUnbundleAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
