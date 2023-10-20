package traveltrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/traveltrade"
)

// Alitriptraveltradequery 飞猪度假-订单详情查询接口
// alitrip.travel.trade.query
//
// 飞猪度假订单详情查询接口
func Alitriptraveltradequery(clt *core.SDKClient, req *traveltrade.AlitriptraveltradequeryAPIRequest, session string) (*traveltrade.AlitriptraveltradequeryAPIResponse, error) {
	var resp traveltrade.AlitriptraveltradequeryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
