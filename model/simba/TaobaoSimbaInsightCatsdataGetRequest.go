package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取类目的大盘数据 APIRequest
taobao.simba.insight.catsdata.get

根据类目id获取类目的大盘数据，包括展现指数，点击指数，点击率，本次提供的insight相关的其它接口的都是这种情况。
*/
type TaobaoSimbaInsightCatsdataGetRequest struct {
    model.Params

    // 表示要查询的类目id
    categoryIdList   []string 

    // 开始时间，格式：yyyy-MM-dd
    startDate   string 

    // 查询截止时间，格式：yyyy-MM-dd
    endDate   string 

}

func NewTaobaoSimbaInsightCatsdataGetRequest() *TaobaoSimbaInsightCatsdataGetRequest{
    return &TaobaoSimbaInsightCatsdataGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaInsightCatsdataGetRequest) GetApiMethodName() string {
    return "taobao.simba.insight.catsdata.get"
}

func (r TaobaoSimbaInsightCatsdataGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaInsightCatsdataGetRequest) SetCategoryIdList(categoryIdList []string) error {
    r.categoryIdList = categoryIdList
    r.Set("category_id_list", categoryIdList)
    return nil
}

func (r TaobaoSimbaInsightCatsdataGetRequest) GetCategoryIdList() []string {
    return r.categoryIdList
}

func (r *TaobaoSimbaInsightCatsdataGetRequest) SetStartDate(startDate string) error {
    r.startDate = startDate
    r.Set("start_date", startDate)
    return nil
}

func (r TaobaoSimbaInsightCatsdataGetRequest) GetStartDate() string {
    return r.startDate
}

func (r *TaobaoSimbaInsightCatsdataGetRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

func (r TaobaoSimbaInsightCatsdataGetRequest) GetEndDate() string {
    return r.endDate
}

