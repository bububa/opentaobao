package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Alibabadchainaoxiangitembatchdeleteasync 货品与组合货品删除
// alibaba.dchain.aoxiang.item.batch.delete.async
//
// 货品与组合货品删除
func Alibabadchainaoxiangitembatchdeleteasync(clt *core.SDKClient, req *ascp.AlibabadchainaoxiangitembatchdeleteasyncAPIRequest, session string) (*ascp.AlibabadchainaoxiangitembatchdeleteasyncAPIResponse, error) {
	var resp ascp.AlibabadchainaoxiangitembatchdeleteasyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
