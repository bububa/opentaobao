package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
超级推荐【商品推广】创意分时报表查询 APIRequest
taobao.feedflow.item.creative.rpthourlist

创意分时数据查询，支持广告主查询最近90天内某一天的创意维度分时报表数据
*/
type TaobaoFeedflowItemCreativeRpthourlistRequest struct {
    model.Params

    // 查询参数
    rptQuery   *RptQueryDto 

}

func NewTaobaoFeedflowItemCreativeRpthourlistRequest() *TaobaoFeedflowItemCreativeRpthourlistRequest{
    return &TaobaoFeedflowItemCreativeRpthourlistRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFeedflowItemCreativeRpthourlistRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.creative.rpthourlist"
}

func (r TaobaoFeedflowItemCreativeRpthourlistRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFeedflowItemCreativeRpthourlistRequest) SetRptQuery(rptQuery *RptQueryDto) error {
    r.rptQuery = rptQuery
    r.Set("rpt_query", rptQuery)
    return nil
}

func (r TaobaoFeedflowItemCreativeRpthourlistRequest) GetRptQuery() *RptQueryDto {
    return r.rptQuery
}

