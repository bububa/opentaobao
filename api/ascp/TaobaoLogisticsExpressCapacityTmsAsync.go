package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// TaobaoLogisticsExpressCapacityTmsAsync 上门取退产能信息同步/更新
// taobao.logistics.express.capacity.tms.async
//
// 上门取退产能信息同步/更新
func TaobaoLogisticsExpressCapacityTmsAsync(clt *core.SDKClient, req *ascp.TaobaoLogisticsExpressCapacityTmsAsyncAPIRequest, resp *ascp.TaobaoLogisticsExpressCapacityTmsAsyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
