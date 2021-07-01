package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaInsightWordspricedataGetAPIRequest
关键词按竞价区间的分布数据 API请求
taobao.simba.insight.wordspricedata.get

获取关键词按竞价区间进行细分的数据 */
type TaobaoSimbaInsightWordspricedataGetAPIRequest struct {
	model.Params
	// 关键词
	_bidword string
	// 开始时间，格式：yyyy-MM-dd
	_startDate string
	// 结束时间，格式：yyyy-MM-dd
	_endDate string
}

// NewTaobaoSimbaInsightWordspricedataGetRequest 初始化TaobaoSimbaInsightWordspricedataGetAPIRequest对象
func NewTaobaoSimbaInsightWordspricedataGetRequest() *TaobaoSimbaInsightWordspricedataGetAPIRequest {
	return &TaobaoSimbaInsightWordspricedataGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaInsightWordspricedataGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.insight.wordspricedata.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaInsightWordspricedataGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Bidword Setter
// 关键词
func (r *TaobaoSimbaInsightWordspricedataGetAPIRequest) SetBidword(_bidword string) error {
	r._bidword = _bidword
	r.Set("bidword", _bidword)
	return nil
}

// Get Bidword Getter
func (r TaobaoSimbaInsightWordspricedataGetAPIRequest) GetBidword() string {
	return r._bidword
}

// Set is StartDate Setter
// 开始时间，格式：yyyy-MM-dd
func (r *TaobaoSimbaInsightWordspricedataGetAPIRequest) SetStartDate(_startDate string) error {
	r._startDate = _startDate
	r.Set("start_date", _startDate)
	return nil
}

// Get StartDate Getter
func (r TaobaoSimbaInsightWordspricedataGetAPIRequest) GetStartDate() string {
	return r._startDate
}

// Set is EndDate Setter
// 结束时间，格式：yyyy-MM-dd
func (r *TaobaoSimbaInsightWordspricedataGetAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// Get EndDate Getter
func (r TaobaoSimbaInsightWordspricedataGetAPIRequest) GetEndDate() string {
	return r._endDate
}
