package traveltrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/traveltrade"
)

// Alitriptraveltradessearch 飞猪度假-订单列表搜索接口
// alitrip.travel.trades.search
//
// 订单列表搜索接口：以订单创建、结束时间、订单状态为搜索条件，搜索过滤出满足条件的卖家订单列表。
func Alitriptraveltradessearch(clt *core.SDKClient, req *traveltrade.AlitriptraveltradessearchAPIRequest, session string) (*traveltrade.AlitriptraveltradessearchAPIResponse, error) {
	var resp traveltrade.AlitriptraveltradessearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
