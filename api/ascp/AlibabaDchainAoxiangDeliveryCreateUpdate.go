package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangDeliveryCreateUpdate 新建/更新配资源
// alibaba.dchain.aoxiang.delivery.create.update
//
// 新建/更新配资源
func AlibabaDchainAoxiangDeliveryCreateUpdate(clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangDeliveryCreateUpdateAPIRequest, session string) (*ascp.AlibabaDchainAoxiangDeliveryCreateUpdateAPIResponse, error) {
	var resp ascp.AlibabaDchainAoxiangDeliveryCreateUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
