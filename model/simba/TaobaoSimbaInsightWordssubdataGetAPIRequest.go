package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbainsightwordssubdatagetAPIRequest 获取关键词按流量细分的数据 API请求
// taobao.simba.insight.wordssubdata.get
//
// 获取关键词按流量进行细分的数据，返回结果中network表示流量的来源，意义如下：1-&gt;PC站内，2-&gt;PC站外,4-&gt;无线站内 5-&gt;无线站外
type TaobaosimbainsightwordssubdatagetAPIRequest struct {
	model.Params
	// 关键词列表
	_bidwordList []string
	// 开始时间
	_startDate string
	// 结束时间
	_endDate string
}

// NewTaobaosimbainsightwordssubdatagetRequest 初始化TaobaosimbainsightwordssubdatagetAPIRequest对象
func NewTaobaosimbainsightwordssubdatagetRequest() *TaobaosimbainsightwordssubdatagetAPIRequest {
	return &TaobaosimbainsightwordssubdatagetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbainsightwordssubdatagetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.insight.wordssubdata.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbainsightwordssubdatagetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbainsightwordssubdatagetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBidwordList is BidwordList Setter
// 关键词列表
func (r *TaobaosimbainsightwordssubdatagetAPIRequest) SetBidwordList(_bidwordList []string) error {
	r._bidwordList = _bidwordList
	r.Set("bidword_list", _bidwordList)
	return nil
}

// GetBidwordList BidwordList Getter
func (r TaobaosimbainsightwordssubdatagetAPIRequest) GetBidwordList() []string {
	return r._bidwordList
}

// SetStartDate is StartDate Setter
// 开始时间
func (r *TaobaosimbainsightwordssubdatagetAPIRequest) SetStartDate(_startDate string) error {
	r._startDate = _startDate
	r.Set("start_date", _startDate)
	return nil
}

// GetStartDate StartDate Getter
func (r TaobaosimbainsightwordssubdatagetAPIRequest) GetStartDate() string {
	return r._startDate
}

// SetEndDate is EndDate Setter
// 结束时间
func (r *TaobaosimbainsightwordssubdatagetAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r TaobaosimbainsightwordssubdatagetAPIRequest) GetEndDate() string {
	return r._endDate
}
