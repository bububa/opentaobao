package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangDeliveryStatusUpdate 启用/停用配资源
// alibaba.dchain.aoxiang.delivery.status.update
//
// 启用/停用配资源
func AlibabaDchainAoxiangDeliveryStatusUpdate(clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangDeliveryStatusUpdateAPIRequest, resp *ascp.AlibabaDchainAoxiangDeliveryStatusUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
