package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaInsightCatsdataGetAPIRequest 获取类目的大盘数据 API请求
// taobao.simba.insight.catsdata.get
//
// 根据类目id获取类目的大盘数据，包括展现指数，点击指数，点击率，本次提供的insight相关的其它接口的都是这种情况。
type TaobaoSimbaInsightCatsdataGetAPIRequest struct {
	model.Params
	// 表示要查询的类目id
	_categoryIdList []string
	// 开始时间，格式：yyyy-MM-dd
	_startDate string
	// 查询截止时间，格式：yyyy-MM-dd
	_endDate string
}

// NewTaobaoSimbaInsightCatsdataGetRequest 初始化TaobaoSimbaInsightCatsdataGetAPIRequest对象
func NewTaobaoSimbaInsightCatsdataGetRequest() *TaobaoSimbaInsightCatsdataGetAPIRequest {
	return &TaobaoSimbaInsightCatsdataGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaInsightCatsdataGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.insight.catsdata.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaInsightCatsdataGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaInsightCatsdataGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCategoryIdList is CategoryIdList Setter
// 表示要查询的类目id
func (r *TaobaoSimbaInsightCatsdataGetAPIRequest) SetCategoryIdList(_categoryIdList []string) error {
	r._categoryIdList = _categoryIdList
	r.Set("category_id_list", _categoryIdList)
	return nil
}

// GetCategoryIdList CategoryIdList Getter
func (r TaobaoSimbaInsightCatsdataGetAPIRequest) GetCategoryIdList() []string {
	return r._categoryIdList
}

// SetStartDate is StartDate Setter
// 开始时间，格式：yyyy-MM-dd
func (r *TaobaoSimbaInsightCatsdataGetAPIRequest) SetStartDate(_startDate string) error {
	r._startDate = _startDate
	r.Set("start_date", _startDate)
	return nil
}

// GetStartDate StartDate Getter
func (r TaobaoSimbaInsightCatsdataGetAPIRequest) GetStartDate() string {
	return r._startDate
}

// SetEndDate is EndDate Setter
// 查询截止时间，格式：yyyy-MM-dd
func (r *TaobaoSimbaInsightCatsdataGetAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r TaobaoSimbaInsightCatsdataGetAPIRequest) GetEndDate() string {
	return r._endDate
}
