package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// TaobaoLogisticsExpressCollectResourceTmsDelete 上门取退可揽范围删除
// taobao.logistics.express.collect.resource.tms.delete
//
// 上门取退可揽范围删除
func TaobaoLogisticsExpressCollectResourceTmsDelete(clt *core.SDKClient, req *ascp.TaobaoLogisticsExpressCollectResourceTmsDeleteAPIRequest, resp *ascp.TaobaoLogisticsExpressCollectResourceTmsDeleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
