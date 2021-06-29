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
type TaobaoSimbaInsightCatsworddataGetRequest struct {
    model.Params
    // 类目id
    catId   string
    // 需要查询的关键词列表，最大长度100。
    bidwordList   []string
    // 开始时间，格式只能为：yyyy-MM-dd
    startDate   string
    // 结束时间，格式只能为：yyyy-MM-dd
    endDate   string
}

// 初始化TaobaoSimbaInsightCatsworddataGetRequest对象
func NewTaobaoSimbaInsightCatsworddataGetRequest() *TaobaoSimbaInsightCatsworddataGetRequest{
    return &TaobaoSimbaInsightCatsworddataGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaInsightCatsworddataGetRequest) GetApiMethodName() string {
    return "taobao.simba.insight.catsworddata.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaInsightCatsworddataGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CatId Setter
// 类目id
func (r *TaobaoSimbaInsightCatsworddataGetRequest) SetCatId(catId string) error {
    r.catId = catId
    r.Set("cat_id", catId)
    return nil
}

// CatId Getter
func (r TaobaoSimbaInsightCatsworddataGetRequest) GetCatId() string {
    return r.catId
}
// BidwordList Setter
// 需要查询的关键词列表，最大长度100。
func (r *TaobaoSimbaInsightCatsworddataGetRequest) SetBidwordList(bidwordList []string) error {
    r.bidwordList = bidwordList
    r.Set("bidword_list", bidwordList)
    return nil
}

// BidwordList Getter
func (r TaobaoSimbaInsightCatsworddataGetRequest) GetBidwordList() []string {
    return r.bidwordList
}
// StartDate Setter
// 开始时间，格式只能为：yyyy-MM-dd
func (r *TaobaoSimbaInsightCatsworddataGetRequest) SetStartDate(startDate string) error {
    r.startDate = startDate
    r.Set("start_date", startDate)
    return nil
}

// StartDate Getter
func (r TaobaoSimbaInsightCatsworddataGetRequest) GetStartDate() string {
    return r.startDate
}
// EndDate Setter
// 结束时间，格式只能为：yyyy-MM-dd
func (r *TaobaoSimbaInsightCatsworddataGetRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

// EndDate Getter
func (r TaobaoSimbaInsightCatsworddataGetRequest) GetEndDate() string {
    return r.endDate
}
