package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaInsightWordsdataGetAPIRequest 获取关键词的大盘数据 API请求
// taobao.simba.insight.wordsdata.get
//
// 获取关键词的详细数据
type TaobaoSimbaInsightWordsdataGetAPIRequest struct {
	model.Params
	// 关键词列表，最多可传100个。
	_bidwordList []string
	// 开始时间
	_startDate string
	// 结束时间
	_endDate string
}

// NewTaobaoSimbaInsightWordsdataGetRequest 初始化TaobaoSimbaInsightWordsdataGetAPIRequest对象
func NewTaobaoSimbaInsightWordsdataGetRequest() *TaobaoSimbaInsightWordsdataGetAPIRequest {
	return &TaobaoSimbaInsightWordsdataGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaInsightWordsdataGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.insight.wordsdata.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaInsightWordsdataGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaInsightWordsdataGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBidwordList is BidwordList Setter
// 关键词列表，最多可传100个。
func (r *TaobaoSimbaInsightWordsdataGetAPIRequest) SetBidwordList(_bidwordList []string) error {
	r._bidwordList = _bidwordList
	r.Set("bidword_list", _bidwordList)
	return nil
}

// GetBidwordList BidwordList Getter
func (r TaobaoSimbaInsightWordsdataGetAPIRequest) GetBidwordList() []string {
	return r._bidwordList
}

// SetStartDate is StartDate Setter
// 开始时间
func (r *TaobaoSimbaInsightWordsdataGetAPIRequest) SetStartDate(_startDate string) error {
	r._startDate = _startDate
	r.Set("start_date", _startDate)
	return nil
}

// GetStartDate StartDate Getter
func (r TaobaoSimbaInsightWordsdataGetAPIRequest) GetStartDate() string {
	return r._startDate
}

// SetEndDate is EndDate Setter
// 结束时间
func (r *TaobaoSimbaInsightWordsdataGetAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r TaobaoSimbaInsightWordsdataGetAPIRequest) GetEndDate() string {
	return r._endDate
}
