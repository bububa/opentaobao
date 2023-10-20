package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Alibabadchainaoxiangconsignordergiftbinding 赠品绑赠计算占用
// alibaba.dchain.aoxiang.consignorder.gift.binding
//
// 赠品绑赠计算占用
func Alibabadchainaoxiangconsignordergiftbinding(clt *core.SDKClient, req *ascp.AlibabadchainaoxiangconsignordergiftbindingAPIRequest, session string) (*ascp.AlibabadchainaoxiangconsignordergiftbindingAPIResponse, error) {
	var resp ascp.AlibabadchainaoxiangconsignordergiftbindingAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
