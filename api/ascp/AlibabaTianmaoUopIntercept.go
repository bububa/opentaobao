package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaTianmaoUopIntercept 阿里巴巴.天猫. 履约订单. 配送拦截
// alibaba.tianmao.uop.intercept
//
// 阿里巴巴.天猫. 履约订单. 配送拦截
func AlibabaTianmaoUopIntercept(clt *core.SDKClient, req *ascp.AlibabaTianmaoUopInterceptAPIRequest, session string) (*ascp.AlibabaTianmaoUopInterceptAPIResponse, error) {
	var resp ascp.AlibabaTianmaoUopInterceptAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
