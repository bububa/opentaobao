package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Alibabadchainaoxiangscitembatchcreate 新建货品
// alibaba.dchain.aoxiang.scitem.batch.create
//
// 新建货品
func Alibabadchainaoxiangscitembatchcreate(clt *core.SDKClient, req *ascp.AlibabadchainaoxiangscitembatchcreateAPIRequest, session string) (*ascp.AlibabadchainaoxiangscitembatchcreateAPIResponse, error) {
	var resp ascp.AlibabadchainaoxiangscitembatchcreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
