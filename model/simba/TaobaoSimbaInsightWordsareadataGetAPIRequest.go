package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaInsightWordsareadataGetAPIRequest 获取关键词按地域进行细分的数据 API请求
// taobao.simba.insight.wordsareadata.get
//
// 获取关键词按地域细分的详细数据，目前地域只能细化到省级别，返回的结果中包括市，是为了方便以后扩展，目前结果中市的值等于省。
type TaobaoSimbaInsightWordsareadataGetAPIRequest struct {
	model.Params
	// 关键词
	_bidword string
	// 开始时间，格式：yyyy-MM-dd
	_startDate string
	// 结束时间，格式：yyyy-MM-dd
	_endDate string
}

// NewTaobaoSimbaInsightWordsareadataGetRequest 初始化TaobaoSimbaInsightWordsareadataGetAPIRequest对象
func NewTaobaoSimbaInsightWordsareadataGetRequest() *TaobaoSimbaInsightWordsareadataGetAPIRequest {
	return &TaobaoSimbaInsightWordsareadataGetAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSimbaInsightWordsareadataGetAPIRequest) Reset() {
	r._bidword = ""
	r._startDate = ""
	r._endDate = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaInsightWordsareadataGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.insight.wordsareadata.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaInsightWordsareadataGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaInsightWordsareadataGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBidword is Bidword Setter
// 关键词
func (r *TaobaoSimbaInsightWordsareadataGetAPIRequest) SetBidword(_bidword string) error {
	r._bidword = _bidword
	r.Set("bidword", _bidword)
	return nil
}

// GetBidword Bidword Getter
func (r TaobaoSimbaInsightWordsareadataGetAPIRequest) GetBidword() string {
	return r._bidword
}

// SetStartDate is StartDate Setter
// 开始时间，格式：yyyy-MM-dd
func (r *TaobaoSimbaInsightWordsareadataGetAPIRequest) SetStartDate(_startDate string) error {
	r._startDate = _startDate
	r.Set("start_date", _startDate)
	return nil
}

// GetStartDate StartDate Getter
func (r TaobaoSimbaInsightWordsareadataGetAPIRequest) GetStartDate() string {
	return r._startDate
}

// SetEndDate is EndDate Setter
// 结束时间，格式：yyyy-MM-dd
func (r *TaobaoSimbaInsightWordsareadataGetAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r TaobaoSimbaInsightWordsareadataGetAPIRequest) GetEndDate() string {
	return r._endDate
}

var poolTaobaoSimbaInsightWordsareadataGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSimbaInsightWordsareadataGetRequest()
	},
}

// GetTaobaoSimbaInsightWordsareadataGetRequest 从 sync.Pool 获取 TaobaoSimbaInsightWordsareadataGetAPIRequest
func GetTaobaoSimbaInsightWordsareadataGetAPIRequest() *TaobaoSimbaInsightWordsareadataGetAPIRequest {
	return poolTaobaoSimbaInsightWordsareadataGetAPIRequest.Get().(*TaobaoSimbaInsightWordsareadataGetAPIRequest)
}

// ReleaseTaobaoSimbaInsightWordsareadataGetAPIRequest 将 TaobaoSimbaInsightWordsareadataGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoSimbaInsightWordsareadataGetAPIRequest(v *TaobaoSimbaInsightWordsareadataGetAPIRequest) {
	v.Reset()
	poolTaobaoSimbaInsightWordsareadataGetAPIRequest.Put(v)
}
