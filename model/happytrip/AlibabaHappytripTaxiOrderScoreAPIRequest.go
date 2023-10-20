package happytrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabahappytriptaxiorderscoreAPIRequest 订单打分和评价 API请求
// alibaba.happytrip.taxi.order.score
//
// 对司机进行评分，只有订单结束后，才能进行。
type AlibabahappytriptaxiorderscoreAPIRequest struct {
	model.Params
	// 订单id
	_orderId string
	// 司机评价最多40个汉字
	_comment string
	// 司机评分 星级(1-5)
	_level int64
}

// NewAlibabahappytriptaxiorderscoreRequest 初始化AlibabahappytriptaxiorderscoreAPIRequest对象
func NewAlibabahappytriptaxiorderscoreRequest() *AlibabahappytriptaxiorderscoreAPIRequest {
	return &AlibabahappytriptaxiorderscoreAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabahappytriptaxiorderscoreAPIRequest) GetApiMethodName() string {
	return "alibaba.happytrip.taxi.order.score"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabahappytriptaxiorderscoreAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabahappytriptaxiorderscoreAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 订单id
func (r *AlibabahappytriptaxiorderscoreAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabahappytriptaxiorderscoreAPIRequest) GetOrderId() string {
	return r._orderId
}

// SetComment is Comment Setter
// 司机评价最多40个汉字
func (r *AlibabahappytriptaxiorderscoreAPIRequest) SetComment(_comment string) error {
	r._comment = _comment
	r.Set("comment", _comment)
	return nil
}

// GetComment Comment Getter
func (r AlibabahappytriptaxiorderscoreAPIRequest) GetComment() string {
	return r._comment
}

// SetLevel is Level Setter
// 司机评分 星级(1-5)
func (r *AlibabahappytriptaxiorderscoreAPIRequest) SetLevel(_level int64) error {
	r._level = _level
	r.Set("level", _level)
	return nil
}

// GetLevel Level Getter
func (r AlibabahappytriptaxiorderscoreAPIRequest) GetLevel() int64 {
	return r._level
}
