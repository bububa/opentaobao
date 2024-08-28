package ascp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangConsignorderCancel 自动流转单据取消仓内发货
// alibaba.dchain.aoxiang.consignorder.cancel
//
// 自动流转单据取消仓内发货
func AlibabaDchainAoxiangConsignorderCancel(ctx context.Context, clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangConsignorderCancelAPIRequest, resp *ascp.AlibabaDchainAoxiangConsignorderCancelAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
