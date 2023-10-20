package happytrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/happytrip"
)

// AlibabaHappytripTaxiOrderScore 订单打分和评价
// alibaba.happytrip.taxi.order.score
//
// 对司机进行评分，只有订单结束后，才能进行。
func AlibabaHappytripTaxiOrderScore(clt *core.SDKClient, req *happytrip.AlibabaHappytripTaxiOrderScoreAPIRequest, resp *happytrip.AlibabaHappytripTaxiOrderScoreAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
