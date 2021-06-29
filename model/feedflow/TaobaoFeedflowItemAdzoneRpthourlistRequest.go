package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
超级推荐【商品推广】资源位分时报表查询 APIRequest
taobao.feedflow.item.adzone.rpthourlist

广告主资源包分时数据查询，支持广告主查询最近90天内某一天的资源包维度分时报表数据
*/
type TaobaoFeedflowItemAdzoneRpthourlistRequest struct {
    model.Params

    // 查询参数
    rptQuery   *RptQueryDto 

}

func NewTaobaoFeedflowItemAdzoneRpthourlistRequest() *TaobaoFeedflowItemAdzoneRpthourlistRequest{
    return &TaobaoFeedflowItemAdzoneRpthourlistRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFeedflowItemAdzoneRpthourlistRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.adzone.rpthourlist"
}

func (r TaobaoFeedflowItemAdzoneRpthourlistRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFeedflowItemAdzoneRpthourlistRequest) SetRptQuery(rptQuery *RptQueryDto) error {
    r.rptQuery = rptQuery
    r.Set("rpt_query", rptQuery)
    return nil
}

func (r TaobaoFeedflowItemAdzoneRpthourlistRequest) GetRptQuery() *RptQueryDto {
    return r.rptQuery
}

