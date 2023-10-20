package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Alibabatianmaouopcancel 阿里巴巴.天猫. 履约订单. 取消
// alibaba.tianmao.uop.cancel
//
// 阿里巴巴.天猫. 履约订单. 取消
func Alibabatianmaouopcancel(clt *core.SDKClient, req *ascp.AlibabatianmaouopcancelAPIRequest, session string) (*ascp.AlibabatianmaouopcancelAPIResponse, error) {
	var resp ascp.AlibabatianmaouopcancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
