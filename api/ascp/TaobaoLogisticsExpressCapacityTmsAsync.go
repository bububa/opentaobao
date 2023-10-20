package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// TaobaoLogisticsExpressCapacityTmsAsync 上门取退产能信息同步/更新
// taobao.logistics.express.capacity.tms.async
//
// 上门取退产能信息同步/更新
func TaobaoLogisticsExpressCapacityTmsAsync(clt *core.SDKClient, req *ascp.TaobaoLogisticsExpressCapacityTmsAsyncAPIRequest, session string) (*ascp.TaobaoLogisticsExpressCapacityTmsAsyncAPIResponse, error) {
	var resp ascp.TaobaoLogisticsExpressCapacityTmsAsyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
