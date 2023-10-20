package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Alibabadchainaoxiangscitembatchupdate alibaba.dchain.aoxiang.scitem.batch.update
// alibaba.dchain.aoxiang.scitem.batch.update
//
// 更新货品
func Alibabadchainaoxiangscitembatchupdate(clt *core.SDKClient, req *ascp.AlibabadchainaoxiangscitembatchupdateAPIRequest, session string) (*ascp.AlibabadchainaoxiangscitembatchupdateAPIResponse, error) {
	var resp ascp.AlibabadchainaoxiangscitembatchupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
