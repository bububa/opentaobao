package logistic

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// TaobaoLogisticsAddressSearch 查询卖家地址库
// taobao.logistics.address.search
//
// 通过此接口查询卖家地址库，
func TaobaoLogisticsAddressSearch(ctx context.Context, clt *core.SDKClient, req *logistic.TaobaoLogisticsAddressSearchAPIRequest, resp *logistic.TaobaoLogisticsAddressSearchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
