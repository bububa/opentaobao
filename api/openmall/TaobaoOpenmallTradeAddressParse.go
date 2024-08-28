package openmall

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openmall"
)

// TaobaoOpenmallTradeAddressParse openmall服务地址区域码解析
// taobao.openmall.trade.address.parse
//
// openmall服务，解析地址区域码，获取创建订单等接口中的区域码信息
func TaobaoOpenmallTradeAddressParse(ctx context.Context, clt *core.SDKClient, req *openmall.TaobaoOpenmallTradeAddressParseAPIRequest, resp *openmall.TaobaoOpenmallTradeAddressParseAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
