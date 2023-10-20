package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Alibabadchainaoxiangitemmappingunbundle 商货关联解绑
// alibaba.dchain.aoxiang.itemmapping.unbundle
//
// 商货关联解绑
func Alibabadchainaoxiangitemmappingunbundle(clt *core.SDKClient, req *ascp.AlibabadchainaoxiangitemmappingunbundleAPIRequest, session string) (*ascp.AlibabadchainaoxiangitemmappingunbundleAPIResponse, error) {
	var resp ascp.AlibabadchainaoxiangitemmappingunbundleAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
