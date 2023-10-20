package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbainsightwordspricedatagetAPIRequest 关键词按竞价区间的分布数据 API请求
// taobao.simba.insight.wordspricedata.get
//
// 获取关键词按竞价区间进行细分的数据
type TaobaosimbainsightwordspricedatagetAPIRequest struct {
	model.Params
	// 关键词
	_bidword string
	// 开始时间，格式：yyyy-MM-dd
	_startDate string
	// 结束时间，格式：yyyy-MM-dd
	_endDate string
}

// NewTaobaosimbainsightwordspricedatagetRequest 初始化TaobaosimbainsightwordspricedatagetAPIRequest对象
func NewTaobaosimbainsightwordspricedatagetRequest() *TaobaosimbainsightwordspricedatagetAPIRequest {
	return &TaobaosimbainsightwordspricedatagetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbainsightwordspricedatagetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.insight.wordspricedata.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbainsightwordspricedatagetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbainsightwordspricedatagetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBidword is Bidword Setter
// 关键词
func (r *TaobaosimbainsightwordspricedatagetAPIRequest) SetBidword(_bidword string) error {
	r._bidword = _bidword
	r.Set("bidword", _bidword)
	return nil
}

// GetBidword Bidword Getter
func (r TaobaosimbainsightwordspricedatagetAPIRequest) GetBidword() string {
	return r._bidword
}

// SetStartDate is StartDate Setter
// 开始时间，格式：yyyy-MM-dd
func (r *TaobaosimbainsightwordspricedatagetAPIRequest) SetStartDate(_startDate string) error {
	r._startDate = _startDate
	r.Set("start_date", _startDate)
	return nil
}

// GetStartDate StartDate Getter
func (r TaobaosimbainsightwordspricedatagetAPIRequest) GetStartDate() string {
	return r._startDate
}

// SetEndDate is EndDate Setter
// 结束时间，格式：yyyy-MM-dd
func (r *TaobaosimbainsightwordspricedatagetAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r TaobaosimbainsightwordspricedatagetAPIRequest) GetEndDate() string {
	return r._endDate
}
