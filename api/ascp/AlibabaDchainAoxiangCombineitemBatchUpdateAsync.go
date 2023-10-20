package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Alibabadchainaoxiangcombineitembatchupdateasync 组合货品新建&更新
// alibaba.dchain.aoxiang.combineitem.batch.update.async
//
// 组合货品新建&amp;更新
func Alibabadchainaoxiangcombineitembatchupdateasync(clt *core.SDKClient, req *ascp.AlibabadchainaoxiangcombineitembatchupdateasyncAPIRequest, session string) (*ascp.AlibabadchainaoxiangcombineitembatchupdateasyncAPIResponse, error) {
	var resp ascp.AlibabadchainaoxiangcombineitembatchupdateasyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
