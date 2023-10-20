package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Alibabadchainaoxiangitembatchupdateasync 货品新建/更新接口
// alibaba.dchain.aoxiang.item.batch.update.async
//
// 货品新建/更新接口
func Alibabadchainaoxiangitembatchupdateasync(clt *core.SDKClient, req *ascp.AlibabadchainaoxiangitembatchupdateasyncAPIRequest, session string) (*ascp.AlibabadchainaoxiangitembatchupdateasyncAPIResponse, error) {
	var resp ascp.AlibabadchainaoxiangitembatchupdateasyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
