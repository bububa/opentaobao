package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取类目下最热门的词 APIRequest
taobao.simba.insight.catstopwordnew.get

按照某个维度，查询某个类目下最热门的词，维度有点击，展现，花费，点击率等，具体可以按哪些字段进行排序，参考参数说明，比如选择了impression，则返回该类目下展现量最高那几个词。
*/
type TaobaoSimbaInsightCatstopwordnewGetRequest struct {
    model.Params

    // 类目id
    catId   string 

    // 查询开始时间，格式必须为：yyyy-MM-dd
    startDate   string 

    // 查询截止时间，格式只能是：yyyy-MM-dd
    endDate   string 

    // 表示查询的维度，比如选择click，则查询该类目下点击量最大的词，可供选择的值有：impression, click, cost, ctr, cpc, coverage, transactiontotal, transactionshippingtotal, favtotal, roi
    dimension   string 

    // 返回前多少条数据
    pageSize   int64 

}

func NewTaobaoSimbaInsightCatstopwordnewGetRequest() *TaobaoSimbaInsightCatstopwordnewGetRequest{
    return &TaobaoSimbaInsightCatstopwordnewGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaInsightCatstopwordnewGetRequest) GetApiMethodName() string {
    return "taobao.simba.insight.catstopwordnew.get"
}

func (r TaobaoSimbaInsightCatstopwordnewGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaInsightCatstopwordnewGetRequest) SetCatId(catId string) error {
    r.catId = catId
    r.Set("cat_id", catId)
    return nil
}

func (r TaobaoSimbaInsightCatstopwordnewGetRequest) GetCatId() string {
    return r.catId
}

func (r *TaobaoSimbaInsightCatstopwordnewGetRequest) SetStartDate(startDate string) error {
    r.startDate = startDate
    r.Set("start_date", startDate)
    return nil
}

func (r TaobaoSimbaInsightCatstopwordnewGetRequest) GetStartDate() string {
    return r.startDate
}

func (r *TaobaoSimbaInsightCatstopwordnewGetRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

func (r TaobaoSimbaInsightCatstopwordnewGetRequest) GetEndDate() string {
    return r.endDate
}

func (r *TaobaoSimbaInsightCatstopwordnewGetRequest) SetDimension(dimension string) error {
    r.dimension = dimension
    r.Set("dimension", dimension)
    return nil
}

func (r TaobaoSimbaInsightCatstopwordnewGetRequest) GetDimension() string {
    return r.dimension
}

func (r *TaobaoSimbaInsightCatstopwordnewGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoSimbaInsightCatstopwordnewGetRequest) GetPageSize() int64 {
    return r.pageSize
}

