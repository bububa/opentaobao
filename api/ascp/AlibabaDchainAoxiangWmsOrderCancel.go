package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangWmsOrderCancel 回传发货单取消通知
// alibaba.dchain.aoxiang.wms.order.cancel
//
// 回传发货单取消通知
func AlibabaDchainAoxiangWmsOrderCancel(clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangWmsOrderCancelAPIRequest, session string) (*ascp.AlibabaDchainAoxiangWmsOrderCancelAPIResponse, error) {
	var resp ascp.AlibabaDchainAoxiangWmsOrderCancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
