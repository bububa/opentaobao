package ascp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangConsignorderGiftBinding 赠品绑赠计算占用
// alibaba.dchain.aoxiang.consignorder.gift.binding
//
// 赠品绑赠计算占用
func AlibabaDchainAoxiangConsignorderGiftBinding(ctx context.Context, clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangConsignorderGiftBindingAPIRequest, resp *ascp.AlibabaDchainAoxiangConsignorderGiftBindingAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
