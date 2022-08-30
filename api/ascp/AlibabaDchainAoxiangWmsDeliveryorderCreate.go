package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangWmsDeliveryorderCreate 回传仓库接单通知
// alibaba.dchain.aoxiang.wms.deliveryorder.create
//
// WMS上报仓库接单节点状态信息，代表接单环节。
func AlibabaDchainAoxiangWmsDeliveryorderCreate(clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangWmsDeliveryorderCreateAPIRequest, session string) (*ascp.AlibabaDchainAoxiangWmsDeliveryorderCreateAPIResponse, error) {
	var resp ascp.AlibabaDchainAoxiangWmsDeliveryorderCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
