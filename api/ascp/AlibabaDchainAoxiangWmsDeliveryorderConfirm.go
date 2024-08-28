package ascp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangWmsDeliveryorderConfirm 回传发货单确认
// alibaba.dchain.aoxiang.wms.deliveryorder.confirm
//
// 回传发货单确认
func AlibabaDchainAoxiangWmsDeliveryorderConfirm(ctx context.Context, clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangWmsDeliveryorderConfirmAPIRequest, resp *ascp.AlibabaDchainAoxiangWmsDeliveryorderConfirmAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
