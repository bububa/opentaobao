package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Alibabadchainaoxiangconsignordergiftreverse 赠品绑赠回滚
// alibaba.dchain.aoxiang.consignorder.gift.reverse
//
// 赠品绑赠回滚
func Alibabadchainaoxiangconsignordergiftreverse(clt *core.SDKClient, req *ascp.AlibabadchainaoxiangconsignordergiftreverseAPIRequest, session string) (*ascp.AlibabadchainaoxiangconsignordergiftreverseAPIResponse, error) {
	var resp ascp.AlibabadchainaoxiangconsignordergiftreverseAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
