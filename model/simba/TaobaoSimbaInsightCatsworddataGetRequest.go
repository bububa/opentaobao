package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取类目下关键词的数据 APIRequest
taobao.simba.insight.catsworddata.get

获取给定词在给定类目下的详细数据
*/
type TaobaoSimbaInsightCatsworddataGetRequest struct {
    model.Params

    // 类目id
    catId   string 

    // 需要查询的关键词列表，最大长度100。
    bidwordList   []String 

    // 开始时间，格式只能为：yyyy-MM-dd
    startDate   string 

    // 结束时间，格式只能为：yyyy-MM-dd
    endDate   string 

}

func NewTaobaoSimbaInsightCatsworddataGetRequest() *TaobaoSimbaInsightCatsworddataGetRequest{
    return &TaobaoSimbaInsightCatsworddataGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaInsightCatsworddataGetRequest) GetApiMethodName() string {
    return "taobao.simba.insight.catsworddata.get"
}

func (r TaobaoSimbaInsightCatsworddataGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaInsightCatsworddataGetRequest) SetCatId(catId string) error {
    r.catId = catId
    r.Set("cat_id", catId)
    return nil
}

func (r TaobaoSimbaInsightCatsworddataGetRequest) GetCatId() string {
    return r.catId
}

func (r *TaobaoSimbaInsightCatsworddataGetRequest) SetBidwordList(bidwordList []String) error {
    r.bidwordList = bidwordList
    r.Set("bidword_list", bidwordList)
    return nil
}

func (r TaobaoSimbaInsightCatsworddataGetRequest) GetBidwordList() []String {
    return r.bidwordList
}

func (r *TaobaoSimbaInsightCatsworddataGetRequest) SetStartDate(startDate string) error {
    r.startDate = startDate
    r.Set("start_date", startDate)
    return nil
}

func (r TaobaoSimbaInsightCatsworddataGetRequest) GetStartDate() string {
    return r.startDate
}

func (r *TaobaoSimbaInsightCatsworddataGetRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

func (r TaobaoSimbaInsightCatsworddataGetRequest) GetEndDate() string {
    return r.endDate
}

