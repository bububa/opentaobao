package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取类目下关键词的数据 API请求
taobao.simba.insight.catsworddata.get

获取给定词在给定类目下的详细数据
*/
type TaobaoSimbaInsightCatsworddataGetAPIRequest struct {
    model.Params
    // 类目id
    _catId   string
    // 需要查询的关键词列表，最大长度100。
    _bidwordList   []string
    // 开始时间，格式只能为：yyyy-MM-dd
    _startDate   string
    // 结束时间，格式只能为：yyyy-MM-dd
    _endDate   string
}

// 初始化TaobaoSimbaInsightCatsworddataGetAPIRequest对象
func NewTaobaoSimbaInsightCatsworddataGetRequest() *TaobaoSimbaInsightCatsworddataGetAPIRequest{
    return &TaobaoSimbaInsightCatsworddataGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaInsightCatsworddataGetAPIRequest) GetApiMethodName() string {
    return "taobao.simba.insight.catsworddata.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaInsightCatsworddataGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CatId Setter
// 类目id
func (r *TaobaoSimbaInsightCatsworddataGetAPIRequest) SetCatId(_catId string) error {
    r._catId = _catId
    r.Set("cat_id", _catId)
    return nil
}

// CatId Getter
func (r TaobaoSimbaInsightCatsworddataGetAPIRequest) GetCatId() string {
    return r._catId
}
// BidwordList Setter
// 需要查询的关键词列表，最大长度100。
func (r *TaobaoSimbaInsightCatsworddataGetAPIRequest) SetBidwordList(_bidwordList []string) error {
    r._bidwordList = _bidwordList
    r.Set("bidword_list", _bidwordList)
    return nil
}

// BidwordList Getter
func (r TaobaoSimbaInsightCatsworddataGetAPIRequest) GetBidwordList() []string {
    return r._bidwordList
}
// StartDate Setter
// 开始时间，格式只能为：yyyy-MM-dd
func (r *TaobaoSimbaInsightCatsworddataGetAPIRequest) SetStartDate(_startDate string) error {
    r._startDate = _startDate
    r.Set("start_date", _startDate)
    return nil
}

// StartDate Getter
func (r TaobaoSimbaInsightCatsworddataGetAPIRequest) GetStartDate() string {
    return r._startDate
}
// EndDate Setter
// 结束时间，格式只能为：yyyy-MM-dd
func (r *TaobaoSimbaInsightCatsworddataGetAPIRequest) SetEndDate(_endDate string) error {
    r._endDate = _endDate
    r.Set("end_date", _endDate)
    return nil
}

// EndDate Getter
func (r TaobaoSimbaInsightCatsworddataGetAPIRequest) GetEndDate() string {
    return r._endDate
}
