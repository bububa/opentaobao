package jipiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitripsellerrefundorderlistfetchAPIRequest 【机票代理商】——退票订单列表提取 API请求
// taobao.alitrip.seller.refundorderlist.fetch
//
// 代理商纬度退票订单列表提取
type TaobaoalitripsellerrefundorderlistfetchAPIRequest struct {
	model.Params
	// 提取数据的结束时间
	_endDate string
	// 提取数据的开始时间
	_startDate string
	// 1：初始，2：接受，3：确认，4：失败，5：买家处理，6：卖家处理，7：等待小二回填，8：退款成功
	_status int64
}

// NewTaobaoalitripsellerrefundorderlistfetchRequest 初始化TaobaoalitripsellerrefundorderlistfetchAPIRequest对象
func NewTaobaoalitripsellerrefundorderlistfetchRequest() *TaobaoalitripsellerrefundorderlistfetchAPIRequest {
	return &TaobaoalitripsellerrefundorderlistfetchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitripsellerrefundorderlistfetchAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.seller.refundorderlist.fetch"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitripsellerrefundorderlistfetchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitripsellerrefundorderlistfetchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEndDate is EndDate Setter
// 提取数据的结束时间
func (r *TaobaoalitripsellerrefundorderlistfetchAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r TaobaoalitripsellerrefundorderlistfetchAPIRequest) GetEndDate() string {
	return r._endDate
}

// SetStartDate is StartDate Setter
// 提取数据的开始时间
func (r *TaobaoalitripsellerrefundorderlistfetchAPIRequest) SetStartDate(_startDate string) error {
	r._startDate = _startDate
	r.Set("start_date", _startDate)
	return nil
}

// GetStartDate StartDate Getter
func (r TaobaoalitripsellerrefundorderlistfetchAPIRequest) GetStartDate() string {
	return r._startDate
}

// SetStatus is Status Setter
// 1：初始，2：接受，3：确认，4：失败，5：买家处理，6：卖家处理，7：等待小二回填，8：退款成功
func (r *TaobaoalitripsellerrefundorderlistfetchAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaoalitripsellerrefundorderlistfetchAPIRequest) GetStatus() int64 {
	return r._status
}
