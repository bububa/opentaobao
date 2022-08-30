package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangScitemQuery 货品查询
// alibaba.dchain.aoxiang.scitem.query
//
// 货品查询
func AlibabaDchainAoxiangScitemQuery(clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangScitemQueryAPIRequest, session string) (*ascp.AlibabaDchainAoxiangScitemQueryAPIResponse, error) {
	var resp ascp.AlibabaDchainAoxiangScitemQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
