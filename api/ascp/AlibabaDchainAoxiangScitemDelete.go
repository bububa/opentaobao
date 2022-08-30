package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangScitemDelete 货品删除
// alibaba.dchain.aoxiang.scitem.delete
//
// 货品删除
func AlibabaDchainAoxiangScitemDelete(clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangScitemDeleteAPIRequest, session string) (*ascp.AlibabaDchainAoxiangScitemDeleteAPIResponse, error) {
	var resp ascp.AlibabaDchainAoxiangScitemDeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
