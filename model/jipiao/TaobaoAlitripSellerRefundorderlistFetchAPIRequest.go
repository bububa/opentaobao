package jipiao

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripSellerRefundorderlistFetchAPIRequest 【机票代理商】——退票订单列表提取 API请求
// taobao.alitrip.seller.refundorderlist.fetch
//
// 代理商纬度退票订单列表提取
type TaobaoAlitripSellerRefundorderlistFetchAPIRequest struct {
	model.Params
	// 提取数据的结束时间
	_endDate string
	// 提取数据的开始时间
	_startDate string
	// 1：初始，2：接受，3：确认，4：失败，5：买家处理，6：卖家处理，7：等待小二回填，8：退款成功
	_status int64
}

// NewTaobaoAlitripSellerRefundorderlistFetchRequest 初始化TaobaoAlitripSellerRefundorderlistFetchAPIRequest对象
func NewTaobaoAlitripSellerRefundorderlistFetchRequest() *TaobaoAlitripSellerRefundorderlistFetchAPIRequest {
	return &TaobaoAlitripSellerRefundorderlistFetchAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripSellerRefundorderlistFetchAPIRequest) Reset() {
	r._endDate = ""
	r._startDate = ""
	r._status = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripSellerRefundorderlistFetchAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.seller.refundorderlist.fetch"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripSellerRefundorderlistFetchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripSellerRefundorderlistFetchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEndDate is EndDate Setter
// 提取数据的结束时间
func (r *TaobaoAlitripSellerRefundorderlistFetchAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r TaobaoAlitripSellerRefundorderlistFetchAPIRequest) GetEndDate() string {
	return r._endDate
}

// SetStartDate is StartDate Setter
// 提取数据的开始时间
func (r *TaobaoAlitripSellerRefundorderlistFetchAPIRequest) SetStartDate(_startDate string) error {
	r._startDate = _startDate
	r.Set("start_date", _startDate)
	return nil
}

// GetStartDate StartDate Getter
func (r TaobaoAlitripSellerRefundorderlistFetchAPIRequest) GetStartDate() string {
	return r._startDate
}

// SetStatus is Status Setter
// 1：初始，2：接受，3：确认，4：失败，5：买家处理，6：卖家处理，7：等待小二回填，8：退款成功
func (r *TaobaoAlitripSellerRefundorderlistFetchAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaoAlitripSellerRefundorderlistFetchAPIRequest) GetStatus() int64 {
	return r._status
}

var poolTaobaoAlitripSellerRefundorderlistFetchAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripSellerRefundorderlistFetchRequest()
	},
}

// GetTaobaoAlitripSellerRefundorderlistFetchRequest 从 sync.Pool 获取 TaobaoAlitripSellerRefundorderlistFetchAPIRequest
func GetTaobaoAlitripSellerRefundorderlistFetchAPIRequest() *TaobaoAlitripSellerRefundorderlistFetchAPIRequest {
	return poolTaobaoAlitripSellerRefundorderlistFetchAPIRequest.Get().(*TaobaoAlitripSellerRefundorderlistFetchAPIRequest)
}

// ReleaseTaobaoAlitripSellerRefundorderlistFetchAPIRequest 将 TaobaoAlitripSellerRefundorderlistFetchAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripSellerRefundorderlistFetchAPIRequest(v *TaobaoAlitripSellerRefundorderlistFetchAPIRequest) {
	v.Reset()
	poolTaobaoAlitripSellerRefundorderlistFetchAPIRequest.Put(v)
}
