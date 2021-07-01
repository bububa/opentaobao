package jipiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripSellerRefundSearchAPIRequest
【机票代理商】退票申请单列表 API请求
taobao.alitrip.seller.refund.search

查询退票申请单列表 */
type TaobaoAlitripSellerRefundSearchAPIRequest struct {
	model.Params
	// 结束时间
	_endTime string
	// 开始时间
	_startTime string
	// 申请单状态（如果为空查询所有状态，1初始 2接受 3确认 4失败 5买家处理 6卖家处理 7等待小二回填 8退款成功）
	_status int64
}

// NewTaobaoAlitripSellerRefundSearchRequest 初始化TaobaoAlitripSellerRefundSearchAPIRequest对象
func NewTaobaoAlitripSellerRefundSearchRequest() *TaobaoAlitripSellerRefundSearchAPIRequest {
	return &TaobaoAlitripSellerRefundSearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripSellerRefundSearchAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.seller.refund.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripSellerRefundSearchAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is EndTime Setter
// 结束时间
func (r *TaobaoAlitripSellerRefundSearchAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// Get EndTime Getter
func (r TaobaoAlitripSellerRefundSearchAPIRequest) GetEndTime() string {
	return r._endTime
}

// Set is StartTime Setter
// 开始时间
func (r *TaobaoAlitripSellerRefundSearchAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// Get StartTime Getter
func (r TaobaoAlitripSellerRefundSearchAPIRequest) GetStartTime() string {
	return r._startTime
}

// Set is Status Setter
// 申请单状态（如果为空查询所有状态，1初始 2接受 3确认 4失败 5买家处理 6卖家处理 7等待小二回填 8退款成功）
func (r *TaobaoAlitripSellerRefundSearchAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// Get Status Getter
func (r TaobaoAlitripSellerRefundSearchAPIRequest) GetStatus() int64 {
	return r._status
}
