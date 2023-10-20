package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangItemmappingUnbundle 商货关联解绑
// alibaba.dchain.aoxiang.itemmapping.unbundle
//
// 商货关联解绑
func AlibabaDchainAoxiangItemmappingUnbundle(clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangItemmappingUnbundleAPIRequest, resp *ascp.AlibabaDchainAoxiangItemmappingUnbundleAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
