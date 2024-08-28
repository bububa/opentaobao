package ascp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// TaobaoLogisticsExpressAddressBlacklistTmsDelete 上门取退可揽范围黑名单删除接口
// taobao.logistics.express.address.blacklist.tms.delete
//
// 上门取退可揽范围黑名单删除接口
func TaobaoLogisticsExpressAddressBlacklistTmsDelete(ctx context.Context, clt *core.SDKClient, req *ascp.TaobaoLogisticsExpressAddressBlacklistTmsDeleteAPIRequest, resp *ascp.TaobaoLogisticsExpressAddressBlacklistTmsDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
