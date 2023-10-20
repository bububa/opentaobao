package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Alibabadchainaoxiangconsignordercancel 自动流转单据取消仓内发货
// alibaba.dchain.aoxiang.consignorder.cancel
//
// 自动流转单据取消仓内发货
func Alibabadchainaoxiangconsignordercancel(clt *core.SDKClient, req *ascp.AlibabadchainaoxiangconsignordercancelAPIRequest, session string) (*ascp.AlibabadchainaoxiangconsignordercancelAPIResponse, error) {
	var resp ascp.AlibabadchainaoxiangconsignordercancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
