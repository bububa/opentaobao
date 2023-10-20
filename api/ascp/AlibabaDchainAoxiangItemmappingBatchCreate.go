package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Alibabadchainaoxiangitemmappingbatchcreate 新建商货品关联
// alibaba.dchain.aoxiang.itemmapping.batch.create
//
// 新建商货品关联
func Alibabadchainaoxiangitemmappingbatchcreate(clt *core.SDKClient, req *ascp.AlibabadchainaoxiangitemmappingbatchcreateAPIRequest, session string) (*ascp.AlibabadchainaoxiangitemmappingbatchcreateAPIResponse, error) {
	var resp ascp.AlibabadchainaoxiangitemmappingbatchcreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
