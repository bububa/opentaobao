package vaccin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/vaccin"
)

// Alibabaalihealthvaccinetradeorderchannelget 通过订单ID与卖家ID获取订单渠道
// alibaba.alihealth.vaccine.trade.order.channel.get
//
// 通过订单ID与卖家ID获取订单渠道
func Alibabaalihealthvaccinetradeorderchannelget(clt *core.SDKClient, req *vaccin.AlibabaalihealthvaccinetradeorderchannelgetAPIRequest, session string) (*vaccin.AlibabaalihealthvaccinetradeorderchannelgetAPIResponse, error) {
	var resp vaccin.AlibabaalihealthvaccinetradeorderchannelgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
