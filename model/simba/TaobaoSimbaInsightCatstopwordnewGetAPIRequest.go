package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaInsightCatstopwordnewGetAPIRequest 获取类目下最热门的词 API请求
// taobao.simba.insight.catstopwordnew.get
//
// 按照某个维度，查询某个类目下最热门的词，维度有点击，展现，花费，点击率等，具体可以按哪些字段进行排序，参考参数说明，比如选择了impression，则返回该类目下展现量最高那几个词。
type TaobaoSimbaInsightCatstopwordnewGetAPIRequest struct {
	model.Params
	// 类目id
	_catId string
	// 查询开始时间，格式必须为：yyyy-MM-dd
	_startDate string
	// 查询截止时间，格式只能是：yyyy-MM-dd
	_endDate string
	// 表示查询的维度，比如选择click，则查询该类目下点击量最大的词，可供选择的值有：impression, click, cost, ctr, cpc, coverage, transactiontotal, transactionshippingtotal, favtotal, roi
	_dimension string
	// 返回前多少条数据
	_pageSize int64
}

// NewTaobaoSimbaInsightCatstopwordnewGetRequest 初始化TaobaoSimbaInsightCatstopwordnewGetAPIRequest对象
func NewTaobaoSimbaInsightCatstopwordnewGetRequest() *TaobaoSimbaInsightCatstopwordnewGetAPIRequest {
	return &TaobaoSimbaInsightCatstopwordnewGetAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSimbaInsightCatstopwordnewGetAPIRequest) Reset() {
	r._catId = ""
	r._startDate = ""
	r._endDate = ""
	r._dimension = ""
	r._pageSize = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaInsightCatstopwordnewGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.insight.catstopwordnew.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaInsightCatstopwordnewGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaInsightCatstopwordnewGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCatId is CatId Setter
// 类目id
func (r *TaobaoSimbaInsightCatstopwordnewGetAPIRequest) SetCatId(_catId string) error {
	r._catId = _catId
	r.Set("cat_id", _catId)
	return nil
}

// GetCatId CatId Getter
func (r TaobaoSimbaInsightCatstopwordnewGetAPIRequest) GetCatId() string {
	return r._catId
}

// SetStartDate is StartDate Setter
// 查询开始时间，格式必须为：yyyy-MM-dd
func (r *TaobaoSimbaInsightCatstopwordnewGetAPIRequest) SetStartDate(_startDate string) error {
	r._startDate = _startDate
	r.Set("start_date", _startDate)
	return nil
}

// GetStartDate StartDate Getter
func (r TaobaoSimbaInsightCatstopwordnewGetAPIRequest) GetStartDate() string {
	return r._startDate
}

// SetEndDate is EndDate Setter
// 查询截止时间，格式只能是：yyyy-MM-dd
func (r *TaobaoSimbaInsightCatstopwordnewGetAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r TaobaoSimbaInsightCatstopwordnewGetAPIRequest) GetEndDate() string {
	return r._endDate
}

// SetDimension is Dimension Setter
// 表示查询的维度，比如选择click，则查询该类目下点击量最大的词，可供选择的值有：impression, click, cost, ctr, cpc, coverage, transactiontotal, transactionshippingtotal, favtotal, roi
func (r *TaobaoSimbaInsightCatstopwordnewGetAPIRequest) SetDimension(_dimension string) error {
	r._dimension = _dimension
	r.Set("dimension", _dimension)
	return nil
}

// GetDimension Dimension Getter
func (r TaobaoSimbaInsightCatstopwordnewGetAPIRequest) GetDimension() string {
	return r._dimension
}

// SetPageSize is PageSize Setter
// 返回前多少条数据
func (r *TaobaoSimbaInsightCatstopwordnewGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoSimbaInsightCatstopwordnewGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

var poolTaobaoSimbaInsightCatstopwordnewGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSimbaInsightCatstopwordnewGetRequest()
	},
}

// GetTaobaoSimbaInsightCatstopwordnewGetRequest 从 sync.Pool 获取 TaobaoSimbaInsightCatstopwordnewGetAPIRequest
func GetTaobaoSimbaInsightCatstopwordnewGetAPIRequest() *TaobaoSimbaInsightCatstopwordnewGetAPIRequest {
	return poolTaobaoSimbaInsightCatstopwordnewGetAPIRequest.Get().(*TaobaoSimbaInsightCatstopwordnewGetAPIRequest)
}

// ReleaseTaobaoSimbaInsightCatstopwordnewGetAPIRequest 将 TaobaoSimbaInsightCatstopwordnewGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoSimbaInsightCatstopwordnewGetAPIRequest(v *TaobaoSimbaInsightCatstopwordnewGetAPIRequest) {
	v.Reset()
	poolTaobaoSimbaInsightCatstopwordnewGetAPIRequest.Put(v)
}
