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

// NewAlibabaHappytripTaxiOrderScoreRequest 初始化AlibabaHappytripTaxiOrderScoreAPIRequest对象
func NewAlibabaHappytripTaxiOrderScoreRequest() *AlibabaHappytripTaxiOrderScoreAPIRequest {
	return &AlibabaHappytripTaxiOrderScoreAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHappytripTaxiOrderScoreAPIRequest) GetApiMethodName() string {
	return "alibaba.happytrip.taxi.order.score"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHappytripTaxiOrderScoreAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OrderId Setter
// 订单id
func (r *AlibabaHappytripTaxiOrderScoreAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// Get OrderId Getter
func (r AlibabaHappytripTaxiOrderScoreAPIRequest) GetOrderId() string {
	return r._orderId
}

// Set is Level Setter
// 司机评分 星级(1-5)
func (r *AlibabaHappytripTaxiOrderScoreAPIRequest) SetLevel(_level int64) error {
	r._level = _level
	r.Set("level", _level)
	return nil
}

// Get Level Getter
func (r AlibabaHappytripTaxiOrderScoreAPIRequest) GetLevel() int64 {
	return r._level
}

// Set is Comment Setter
// 司机评价最多40个汉字
func (r *AlibabaHappytripTaxiOrderScoreAPIRequest) SetComment(_comment string) error {
	r._comment = _comment
	r.Set("comment", _comment)
	return nil
}

// Get Comment Getter
func (r AlibabaHappytripTaxiOrderScoreAPIRequest) GetComment() string {
	return r._comment
}
