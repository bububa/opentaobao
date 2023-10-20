package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangScitemBatchCreate 新建货品
// alibaba.dchain.aoxiang.scitem.batch.create
//
// 新建货品
func AlibabaDchainAoxiangScitemBatchCreate(clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangScitemBatchCreateAPIRequest, resp *ascp.AlibabaDchainAoxiangScitemBatchCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
