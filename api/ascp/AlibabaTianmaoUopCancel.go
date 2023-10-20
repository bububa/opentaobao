package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaTianmaoUopCancel 阿里巴巴.天猫. 履约订单. 取消
// alibaba.tianmao.uop.cancel
//
// 阿里巴巴.天猫. 履约订单. 取消
func AlibabaTianmaoUopCancel(clt *core.SDKClient, req *ascp.AlibabaTianmaoUopCancelAPIRequest, resp *ascp.AlibabaTianmaoUopCancelAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
