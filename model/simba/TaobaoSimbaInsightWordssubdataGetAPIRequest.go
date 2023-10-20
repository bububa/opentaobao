package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaInsightWordssubdataGetAPIRequest 获取关键词按流量细分的数据 API请求
// taobao.simba.insight.wordssubdata.get
//
// 获取关键词按流量进行细分的数据，返回结果中network表示流量的来源，意义如下：1-&gt;PC站内，2-&gt;PC站外,4-&gt;无线站内 5-&gt;无线站外
type TaobaoSimbaInsightWordssubdataGetAPIRequest struct {
	model.Params
	// 关键词列表
	_bidwordList []string
	// 开始时间
	_startDate string
	// 结束时间
	_endDate string
}

// NewTaobaoSimbaInsightWordssubdataGetRequest 初始化TaobaoSimbaInsightWordssubdataGetAPIRequest对象
func NewTaobaoSimbaInsightWordssubdataGetRequest() *TaobaoSimbaInsightWordssubdataGetAPIRequest {
	return &TaobaoSimbaInsightWordssubdataGetAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSimbaInsightWordssubdataGetAPIRequest) Reset() {
	r._bidwordList = r._bidwordList[:0]
	r._startDate = ""
	r._endDate = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaInsightWordssubdataGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.insight.wordssubdata.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaInsightWordssubdataGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaInsightWordssubdataGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBidwordList is BidwordList Setter
// 关键词列表
func (r *TaobaoSimbaInsightWordssubdataGetAPIRequest) SetBidwordList(_bidwordList []string) error {
	r._bidwordList = _bidwordList
	r.Set("bidword_list", _bidwordList)
	return nil
}

// GetBidwordList BidwordList Getter
func (r TaobaoSimbaInsightWordssubdataGetAPIRequest) GetBidwordList() []string {
	return r._bidwordList
}

// SetStartDate is StartDate Setter
// 开始时间
func (r *TaobaoSimbaInsightWordssubdataGetAPIRequest) SetStartDate(_startDate string) error {
	r._startDate = _startDate
	r.Set("start_date", _startDate)
	return nil
}

// GetStartDate StartDate Getter
func (r TaobaoSimbaInsightWordssubdataGetAPIRequest) GetStartDate() string {
	return r._startDate
}

// SetEndDate is EndDate Setter
// 结束时间
func (r *TaobaoSimbaInsightWordssubdataGetAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r TaobaoSimbaInsightWordssubdataGetAPIRequest) GetEndDate() string {
	return r._endDate
}

var poolTaobaoSimbaInsightWordssubdataGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSimbaInsightWordssubdataGetRequest()
	},
}

// GetTaobaoSimbaInsightWordssubdataGetRequest 从 sync.Pool 获取 TaobaoSimbaInsightWordssubdataGetAPIRequest
func GetTaobaoSimbaInsightWordssubdataGetAPIRequest() *TaobaoSimbaInsightWordssubdataGetAPIRequest {
	return poolTaobaoSimbaInsightWordssubdataGetAPIRequest.Get().(*TaobaoSimbaInsightWordssubdataGetAPIRequest)
}

// ReleaseTaobaoSimbaInsightWordssubdataGetAPIRequest 将 TaobaoSimbaInsightWordssubdataGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoSimbaInsightWordssubdataGetAPIRequest(v *TaobaoSimbaInsightWordssubdataGetAPIRequest) {
	v.Reset()
	poolTaobaoSimbaInsightWordssubdataGetAPIRequest.Put(v)
}
