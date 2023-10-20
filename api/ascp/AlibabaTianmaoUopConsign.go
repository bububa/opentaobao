package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Alibabatianmaouopconsign 阿里巴巴.天猫. 履约订单. 发货
// alibaba.tianmao.uop.consign
//
// 阿里巴巴.天猫. 履约订单. 发货
func Alibabatianmaouopconsign(clt *core.SDKClient, req *ascp.AlibabatianmaouopconsignAPIRequest, session string) (*ascp.AlibabatianmaouopconsignAPIResponse, error) {
	var resp ascp.AlibabatianmaouopconsignAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
