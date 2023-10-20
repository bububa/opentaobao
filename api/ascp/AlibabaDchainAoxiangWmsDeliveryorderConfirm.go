package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangWmsDeliveryorderConfirm 回传发货单确认
// alibaba.dchain.aoxiang.wms.deliveryorder.confirm
//
// 回传发货单确认
func AlibabaDchainAoxiangWmsDeliveryorderConfirm(clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangWmsDeliveryorderConfirmAPIRequest, session string) (*ascp.AlibabaDchainAoxiangWmsDeliveryorderConfirmAPIResponse, error) {
	var resp ascp.AlibabaDchainAoxiangWmsDeliveryorderConfirmAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
