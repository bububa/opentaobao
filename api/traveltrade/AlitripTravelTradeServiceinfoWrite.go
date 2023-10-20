package traveltrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/traveltrade"
)

// AlitripTravelTradeServiceinfoWrite 订单服务信息写入接口
// alitrip.travel.trade.serviceinfo.write
//
// 订单服务信息写入接口
func AlitripTravelTradeServiceinfoWrite(clt *core.SDKClient, req *traveltrade.AlitripTravelTradeServiceinfoWriteAPIRequest, resp *traveltrade.AlitripTravelTradeServiceinfoWriteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
