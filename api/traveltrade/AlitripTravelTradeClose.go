package traveltrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/traveltrade"
)

// Alitriptraveltradeclose 飞猪度假-订单关闭接口（快速退款）
// alitrip.travel.trade.close
//
// 卖家关单（快速退款接口），不支持二次预约商品的订单
func Alitriptraveltradeclose(clt *core.SDKClient, req *traveltrade.AlitriptraveltradecloseAPIRequest, session string) (*traveltrade.AlitriptraveltradecloseAPIResponse, error) {
	var resp traveltrade.AlitriptraveltradecloseAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
