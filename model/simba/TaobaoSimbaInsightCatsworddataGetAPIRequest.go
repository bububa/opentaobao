package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaInsightCatsworddataGetAPIRequest 获取类目下关键词的数据 API请求
// taobao.simba.insight.catsworddata.get
//
// 获取给定词在给定类目下的详细数据
type TaobaoSimbaInsightCatsworddataGetAPIRequest struct {
	model.Params
	// 需要查询的关键词列表，最大长度100。
	_bidwordList []string
	// 类目id
	_catId string
	// 开始时间，格式只能为：yyyy-MM-dd
	_startDate string
	// 结束时间，格式只能为：yyyy-MM-dd
	_endDate string
}

// NewTaobaoSimbaInsightCatsworddataGetRequest 初始化TaobaoSimbaInsightCatsworddataGetAPIRequest对象
func NewTaobaoSimbaInsightCatsworddataGetRequest() *TaobaoSimbaInsightCatsworddataGetAPIRequest {
	return &TaobaoSimbaInsightCatsworddataGetAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSimbaInsightCatsworddataGetAPIRequest) Reset() {
	r._bidwordList = r._bidwordList[:0]
	r._catId = ""
	r._startDate = ""
	r._endDate = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaInsightCatsworddataGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.insight.catsworddata.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaInsightCatsworddataGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaInsightCatsworddataGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBidwordList is BidwordList Setter
// 需要查询的关键词列表，最大长度100。
func (r *TaobaoSimbaInsightCatsworddataGetAPIRequest) SetBidwordList(_bidwordList []string) error {
	r._bidwordList = _bidwordList
	r.Set("bidword_list", _bidwordList)
	return nil
}

// GetBidwordList BidwordList Getter
func (r TaobaoSimbaInsightCatsworddataGetAPIRequest) GetBidwordList() []string {
	return r._bidwordList
}

// SetCatId is CatId Setter
// 类目id
func (r *TaobaoSimbaInsightCatsworddataGetAPIRequest) SetCatId(_catId string) error {
	r._catId = _catId
	r.Set("cat_id", _catId)
	return nil
}

// GetCatId CatId Getter
func (r TaobaoSimbaInsightCatsworddataGetAPIRequest) GetCatId() string {
	return r._catId
}

// SetStartDate is StartDate Setter
// 开始时间，格式只能为：yyyy-MM-dd
func (r *TaobaoSimbaInsightCatsworddataGetAPIRequest) SetStartDate(_startDate string) error {
	r._startDate = _startDate
	r.Set("start_date", _startDate)
	return nil
}

// GetStartDate StartDate Getter
func (r TaobaoSimbaInsightCatsworddataGetAPIRequest) GetStartDate() string {
	return r._startDate
}

// SetEndDate is EndDate Setter
// 结束时间，格式只能为：yyyy-MM-dd
func (r *TaobaoSimbaInsightCatsworddataGetAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r TaobaoSimbaInsightCatsworddataGetAPIRequest) GetEndDate() string {
	return r._endDate
}

var poolTaobaoSimbaInsightCatsworddataGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSimbaInsightCatsworddataGetRequest()
	},
}

// GetTaobaoSimbaInsightCatsworddataGetRequest 从 sync.Pool 获取 TaobaoSimbaInsightCatsworddataGetAPIRequest
func GetTaobaoSimbaInsightCatsworddataGetAPIRequest() *TaobaoSimbaInsightCatsworddataGetAPIRequest {
	return poolTaobaoSimbaInsightCatsworddataGetAPIRequest.Get().(*TaobaoSimbaInsightCatsworddataGetAPIRequest)
}

// ReleaseTaobaoSimbaInsightCatsworddataGetAPIRequest 将 TaobaoSimbaInsightCatsworddataGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoSimbaInsightCatsworddataGetAPIRequest(v *TaobaoSimbaInsightCatsworddataGetAPIRequest) {
	v.Reset()
	poolTaobaoSimbaInsightCatsworddataGetAPIRequest.Put(v)
}
