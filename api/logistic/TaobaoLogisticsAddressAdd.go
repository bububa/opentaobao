package logistic

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// TaobaoLogisticsAddressAdd 卖家地址库新增接口
// taobao.logistics.address.add
//
// 通过此接口新增卖家地址库,卖家最多可添加5条地址库,新增第一条卖家地址，将会自动设为默认地址库
func TaobaoLogisticsAddressAdd(ctx context.Context, clt *core.SDKClient, req *logistic.TaobaoLogisticsAddressAddAPIRequest, resp *logistic.TaobaoLogisticsAddressAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
