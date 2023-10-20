package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Alibabadchainaoxiangcombinescitembatchupdate 更新组合货品
// alibaba.dchain.aoxiang.combinescitem.batch.update
//
// 更新组合货品
func Alibabadchainaoxiangcombinescitembatchupdate(clt *core.SDKClient, req *ascp.AlibabadchainaoxiangcombinescitembatchupdateAPIRequest, session string) (*ascp.AlibabadchainaoxiangcombinescitembatchupdateAPIResponse, error) {
	var resp ascp.AlibabadchainaoxiangcombinescitembatchupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
