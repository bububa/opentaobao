package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangWmsOrderCancel 回传发货单取消通知
// alibaba.dchain.aoxiang.wms.order.cancel
//
// 回传发货单取消通知
func AlibabaDchainAoxiangWmsOrderCancel(clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangWmsOrderCancelAPIRequest, resp *ascp.AlibabaDchainAoxiangWmsOrderCancelAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
