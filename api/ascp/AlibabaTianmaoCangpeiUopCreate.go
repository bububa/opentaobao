package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaTianmaoCangpeiUopCreate 阿里巴巴.天猫家装.仓配.履约订单.创建
// alibaba.tianmao.cangpei.uop.create
//
// 阿里巴巴.天猫家装.仓配.履约订单.创建
func AlibabaTianmaoCangpeiUopCreate(clt *core.SDKClient, req *ascp.AlibabaTianmaoCangpeiUopCreateAPIRequest, resp *ascp.AlibabaTianmaoCangpeiUopCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
