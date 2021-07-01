package happytrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaHappytripTaxiOrderScoreAPIRequest
订单打分和评价 API请求
alibaba.happytrip.taxi.order.score

对司机进行评分，只有订单结束后，才能进行。 */
type AlibabaHappytripTaxiOrderScoreAPIRequest struct {
	model.Params
	// 订单id
	_orderId string
	// 司机评分 星级(1-5)
	_level int64
	// 司机评价最多40个汉字
	_comment string
}

// New
