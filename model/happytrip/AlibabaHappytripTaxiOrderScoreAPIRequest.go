package happytrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHappytripTaxiOrderScoreAPIRequest 订单打分和评价 API请求
// alibaba.happytrip.taxi.order.score
//
// 对司机进行评分，只有订单结束后，才能进行。
type AlibabaHappytripTaxiOrderScoreAPIRequest struct {
	model.Params
	// 订单id
	_orderId string
	// 司机评价最多40个汉字
	_comment string
	// 司机评分 星级(1-5)
	_level int64
}

// NewAlibabaHappytripTaxiOrderScoreRequest 初始化AlibabaHappytripTaxiOrderScoreAPIRequest对象
func NewAlibabaHappytripTaxiOrderScoreRequest() *AlibabaHappytripTaxiOrderScoreAPIRequest {
	return &AlibabaHappytripTaxiOrderScoreAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaHappytripTaxiOrderScoreAPIRequest) Reset() {
	r._orderId = ""
	r._comment = ""
	r._level = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHappytripTaxiOrderScoreAPIRequest) GetApiMethodName() string {
	return "alibaba.happytrip.taxi.order.score"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHappytripTaxiOrderScoreAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHappytripTaxiOrderScoreAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 订单id
func (r *AlibabaHappytripTaxiOrderScoreAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabaHappytripTaxiOrderScoreAPIRequest) GetOrderId() string {
	return r._orderId
}

// SetComment is Comment Setter
// 司机评价最多40个汉字
func (r *AlibabaHappytripTaxiOrderScoreAPIRequest) SetComment(_comment string) error {
	r._comment = _comment
	r.Set("comment", _comment)
	return nil
}

// GetComment Comment Getter
func (r AlibabaHappytripTaxiOrderScoreAPIRequest) GetComment() string {
	return r._comment
}

// SetLevel is Level Setter
// 司机评分 星级(1-5)
func (r *AlibabaHappytripTaxiOrderScoreAPIRequest) SetLevel(_level int64) error {
	r._level = _level
	r.Set("level", _level)
	return nil
}

// GetLevel Level Getter
func (r AlibabaHappytripTaxiOrderScoreAPIRequest) GetLevel() int64 {
	return r._level
}

var poolAlibabaHappytripTaxiOrderScoreAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaHappytripTaxiOrderScoreRequest()
	},
}

// GetAlibabaHappytripTaxiOrderScoreRequest 从 sync.Pool 获取 AlibabaHappytripTaxiOrderScoreAPIRequest
func GetAlibabaHappytripTaxiOrderScoreAPIRequest() *AlibabaHappytripTaxiOrderScoreAPIRequest {
	return poolAlibabaHappytripTaxiOrderScoreAPIRequest.Get().(*AlibabaHappytripTaxiOrderScoreAPIRequest)
}

// ReleaseAlibabaHappytripTaxiOrderScoreAPIRequest 将 AlibabaHappytripTaxiOrderScoreAPIRequest 放入 sync.Pool
func ReleaseAlibabaHappytripTaxiOrderScoreAPIRequest(v *AlibabaHappytripTaxiOrderScoreAPIRequest) {
	v.Reset()
	poolAlibabaHappytripTaxiOrderScoreAPIRequest.Put(v)
}
