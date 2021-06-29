package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取类目的大盘数据 API请求
taobao.simba.insight.catsdata.get

根据类目id获取类目的大盘数据，包括展现指数，点击指数，点击率，本次提供的insight相关的其它接口的都是这种情况。
*/
type TaobaoSimbaInsightCatsdataGetRequest struct {
    model.Params
    // 表示要查询的类目id
    _categoryIdList   []string
    // 开始时间，格式：yyyy-MM-dd
    _startDate   string
    // 查询截止时间，格式：yyyy-MM-dd
    _endDate   string
}

// 初始化TaobaoSimbaInsightCatsdataGetRequest对象
func NewTaobaoSimbaInsightCatsdataGetRequest() *TaobaoSimbaInsightCatsdataGetRequest{
    return &TaobaoSimbaInsightCatsdataGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaInsightCatsdataGetRequest) GetApiMethodName() string {
    return "taobao.simba.insight.catsdata.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaInsightCatsdataGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CategoryIdList Setter
// 表示要查询的类目id
func (r *TaobaoSimbaInsightCatsdataGetRequest) SetCategoryIdList(_categoryIdList []string) error {
    r._categoryIdList = _categoryIdList
    r.Set("category_id_list", _categoryIdList)
    return nil
}

// CategoryIdList Getter
func (r TaobaoSimbaInsightCatsdataGetRequest) GetCategoryIdList() []string {
    return r._categoryIdList
}
// StartDate Setter
// 开始时间，格式：yyyy-MM-dd
func (r *TaobaoSimbaInsightCatsdataGetRequest) SetStartDate(_startDate string) error {
    r._startDate = _startDate
    r.Set("start_date", _startDate)
    return nil
}

// StartDate Getter
func (r TaobaoSimbaInsightCatsdataGetRequest) GetStartDate() string {
    return r._startDate
}
// EndDate Setter
// 查询截止时间，格式：yyyy-MM-dd
func (r *TaobaoSimbaInsightCatsdataGetRequest) SetEndDate(_endDate string) error {
    r._endDate = _endDate
    r.Set("end_date", _endDate)
    return nil
}

// EndDate Getter
func (r TaobaoSimbaInsightCatsdataGetRequest) GetEndDate() string {
    return r._endDate
}
