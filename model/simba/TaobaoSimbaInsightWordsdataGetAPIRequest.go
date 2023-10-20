package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbainsightwordsdatagetAPIRequest 获取关键词的大盘数据 API请求
// taobao.simba.insight.wordsdata.get
//
// 获取关键词的详细数据
type TaobaosimbainsightwordsdatagetAPIRequest struct {
	model.Params
	// 关键词列表，最多可传100个。
	_bidwordList []string
	// 开始时间
	_startDate string
	// 结束时间
	_endDate string
}

// NewTaobaosimbainsightwordsdatagetRequest 初始化TaobaosimbainsightwordsdatagetAPIRequest对象
func NewTaobaosimbainsightwordsdatagetRequest() *TaobaosimbainsightwordsdatagetAPIRequest {
	return &TaobaosimbainsightwordsdatagetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbainsightwordsdatagetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.insight.wordsdata.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbainsightwordsdatagetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbainsightwordsdatagetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBidwordList is BidwordList Setter
// 关键词列表，最多可传100个。
func (r *TaobaosimbainsightwordsdatagetAPIRequest) SetBidwordList(_bidwordList []string) error {
	r._bidwordList = _bidwordList
	r.Set("bidword_list", _bidwordList)
	return nil
}

// GetBidwordList BidwordList Getter
func (r TaobaosimbainsightwordsdatagetAPIRequest) GetBidwordList() []string {
	return r._bidwordList
}

// SetStartDate is StartDate Setter
// 开始时间
func (r *TaobaosimbainsightwordsdatagetAPIRequest) SetStartDate(_startDate string) error {
	r._startDate = _startDate
	r.Set("start_date", _startDate)
	return nil
}

// GetStartDate StartDate Getter
func (r TaobaosimbainsightwordsdatagetAPIRequest) GetStartDate() string {
	return r._startDate
}

// SetEndDate is EndDate Setter
// 结束时间
func (r *TaobaosimbainsightwordsdatagetAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r TaobaosimbainsightwordsdatagetAPIRequest) GetEndDate() string {
	return r._endDate
}
