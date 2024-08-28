package vaccin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/vaccin"
)

// AlibabaAlihealthVaccineTradeOrderChannelGet 通过订单ID与卖家ID获取订单渠道
// alibaba.alihealth.vaccine.trade.order.channel.get
//
// 通过订单ID与卖家ID获取订单渠道
func AlibabaAlihealthVaccineTradeOrderChannelGet(ctx context.Context, clt *core.SDKClient, req *vaccin.AlibabaAlihealthVaccineTradeOrderChannelGetAPIRequest, resp *vaccin.AlibabaAlihealthVaccineTradeOrderChannelGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
