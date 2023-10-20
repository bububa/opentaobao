package vaccin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/vaccin"
)

// AlibabaAlihealthVaccineTradeOrderChannelGet 通过订单ID与卖家ID获取订单渠道
// alibaba.alihealth.vaccine.trade.order.channel.get
//
// 通过订单ID与卖家ID获取订单渠道
func AlibabaAlihealthVaccineTradeOrderChannelGet(clt *core.SDKClient, req *vaccin.AlibabaAlihealthVaccineTradeOrderChannelGetAPIRequest, session string) (*vaccin.AlibabaAlihealthVaccineTradeOrderChannelGetAPIResponse, error) {
	var resp vaccin.AlibabaAlihealthVaccineTradeOrderChannelGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
