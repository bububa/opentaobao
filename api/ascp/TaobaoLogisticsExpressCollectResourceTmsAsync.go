package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// TaobaoLogisticsExpressCollectResourceTmsAsync 配服务商揽收能力同步接口
// taobao.logistics.express.collect.resource.tms.async
//
// 配服务商揽收能力同步接口
func TaobaoLogisticsExpressCollectResourceTmsAsync(clt *core.SDKClient, req *ascp.TaobaoLogisticsExpressCollectResourceTmsAsyncAPIRequest, session string) (*ascp.TaobaoLogisticsExpressCollectResourceTmsAsyncAPIResponse, error) {
	var resp ascp.TaobaoLogisticsExpressCollectResourceTmsAsyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
