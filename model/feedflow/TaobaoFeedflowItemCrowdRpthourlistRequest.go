package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
超级推荐【商品推广】定向分时报表查询 APIRequest
taobao.feedflow.item.crowd.rpthourlist

广告主定向分时数据查询，支持广告主查询最近90天内某一天的定向维度分时报表数据
*/
type TaobaoFeedflowItemCrowdRpthourlistRequest struct {
    model.Params

    // 查询参数
    rptQuery   *RptQueryDto 

}

func NewTaobaoFeedflowItemCrowdRpthourlistRequest() *TaobaoFeedflowItemCrowdRpthourlistRequest{
    return &TaobaoFeedflowItemCrowdRpthourlistRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFeedflowItemCrowdRpthourlistRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.crowd.rpthourlist"
}

func (r TaobaoFeedflowItemCrowdRpthourlistRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFeedflowItemCrowdRpthourlistRequest) SetRptQuery(rptQuery *RptQueryDto) error {
    r.rptQuery = rptQuery
    r.Set("rpt_query", rptQuery)
    return nil
}

func (r TaobaoFeedflowItemCrowdRpthourlistRequest) GetRptQuery() *RptQueryDto {
    return r.rptQuery
}

