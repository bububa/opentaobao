package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Alibabatianmaouopintercept 阿里巴巴.天猫. 履约订单. 配送拦截
// alibaba.tianmao.uop.intercept
//
// 阿里巴巴.天猫. 履约订单. 配送拦截
func Alibabatianmaouopintercept(clt *core.SDKClient, req *ascp.AlibabatianmaouopinterceptAPIRequest, session string) (*ascp.AlibabatianmaouopinterceptAPIResponse, error) {
	var resp ascp.AlibabatianmaouopinterceptAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
