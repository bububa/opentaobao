package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Alibabadchainaoxiangcombinescitembatchcreate 新建组合货品
// alibaba.dchain.aoxiang.combinescitem.batch.create
//
// 新建组合货品
func Alibabadchainaoxiangcombinescitembatchcreate(clt *core.SDKClient, req *ascp.AlibabadchainaoxiangcombinescitembatchcreateAPIRequest, session string) (*ascp.AlibabadchainaoxiangcombinescitembatchcreateAPIResponse, error) {
	var resp ascp.AlibabadchainaoxiangcombinescitembatchcreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
