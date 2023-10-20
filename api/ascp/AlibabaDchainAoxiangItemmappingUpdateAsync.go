package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Alibabadchainaoxiangitemmappingupdateasync 创建/更新商货品关联关系
// alibaba.dchain.aoxiang.itemmapping.update.async
//
// 创建/更新商货品关联关系
func Alibabadchainaoxiangitemmappingupdateasync(clt *core.SDKClient, req *ascp.AlibabadchainaoxiangitemmappingupdateasyncAPIRequest, session string) (*ascp.AlibabadchainaoxiangitemmappingupdateasyncAPIResponse, error) {
	var resp ascp.AlibabadchainaoxiangitemmappingupdateasyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
