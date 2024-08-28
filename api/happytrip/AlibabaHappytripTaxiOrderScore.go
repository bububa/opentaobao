package happytrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/happytrip"
)

// AlibabaHappytripTaxiOrderScore 订单打分和评价
// alibaba.happytrip.taxi.order.score
//
// 对司机进行评分，只有订单结束后，才能进行。
func AlibabaHappytripTaxiOrderScore(ctx context.Context, clt *core.SDKClient, req *happytrip.AlibabaHappytripTaxiOrderScoreAPIRequest, resp *happytrip.AlibabaHappytripTaxiOrderScoreAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
