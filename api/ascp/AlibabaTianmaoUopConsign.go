package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaTianmaoUopConsign 阿里巴巴.天猫. 履约订单. 发货
// alibaba.tianmao.uop.consign
//
// 阿里巴巴.天猫. 履约订单. 发货
func AlibabaTianmaoUopConsign(clt *core.SDKClient, req *ascp.AlibabaTianmaoUopConsignAPIRequest, session string) (*ascp.AlibabaTianmaoUopConsignAPIResponse, error) {
	var resp ascp.AlibabaTianmaoUopConsignAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
