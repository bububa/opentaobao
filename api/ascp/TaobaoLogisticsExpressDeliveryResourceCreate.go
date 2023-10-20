package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// TaobaoLogisticsExpressDeliveryResourceCreate 新建/更新配资源
// taobao.logistics.express.delivery.resource.create
//
// 新建/更新配资源
func TaobaoLogisticsExpressDeliveryResourceCreate(clt *core.SDKClient, req *ascp.TaobaoLogisticsExpressDeliveryResourceCreateAPIRequest, session string) (*ascp.TaobaoLogisticsExpressDeliveryResourceCreateAPIResponse, error) {
	var resp ascp.TaobaoLogisticsExpressDeliveryResourceCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
