package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Alibabadchainaoxiangitemmappingdelete 删除商货品关联关系
// alibaba.dchain.aoxiang.itemmapping.delete
//
// 删除商货品关联关系
func Alibabadchainaoxiangitemmappingdelete(clt *core.SDKClient, req *ascp.AlibabadchainaoxiangitemmappingdeleteAPIRequest, session string) (*ascp.AlibabadchainaoxiangitemmappingdeleteAPIResponse, error) {
	var resp ascp.AlibabadchainaoxiangitemmappingdeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
