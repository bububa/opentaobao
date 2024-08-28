package logistic

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// TaobaoLogisticsAddressRemove 删除卖家地址库
// taobao.logistics.address.remove
//
// 用此接口删除卖家地址库
func TaobaoLogisticsAddressRemove(ctx context.Context, clt *core.SDKClient, req *logistic.TaobaoLogisticsAddressRemoveAPIRequest, resp *logistic.TaobaoLogisticsAddressRemoveAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
