package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// TaobaoLogisticsExpressAddressBlacklistTmsAsync 上门取退可揽范围黑名单同步/更新
// taobao.logistics.express.address.blacklist.tms.async
//
// 上门取退可揽范围黑名单同步/更新
func TaobaoLogisticsExpressAddressBlacklistTmsAsync(clt *core.SDKClient, req *ascp.TaobaoLogisticsExpressAddressBlacklistTmsAsyncAPIRequest, session string) (*ascp.TaobaoLogisticsExpressAddressBlacklistTmsAsyncAPIResponse, error) {
	var resp ascp.TaobaoLogisticsExpressAddressBlacklistTmsAsyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
