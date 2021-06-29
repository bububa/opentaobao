package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
超级推荐广告主分时报表查询 APIRequest
taobao.feedflow.account.rpthourlist

广告主分时报表查询，支持广告主查询最近90天内某一天的账户维度分时报表数据
*/
type TaobaoFeedflowAccountRpthourlistRequest struct {
    model.Params

    // 查询参数
    rptQuery   *RptQueryDto 

}

func NewTaobaoFeedflowAccountRpthourlistRequest() *TaobaoFeedflowAccountRpthourlistRequest{
    return &TaobaoFeedflowAccountRpthourlistRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFeedflowAccountRpthourlistRequest) GetApiMethodName() string {
    return "taobao.feedflow.account.rpthourlist"
}

func (r TaobaoFeedflowAccountRpthourlistRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFeedflowAccountRpthourlistRequest) SetRptQuery(rptQuery *RptQueryDto) error {
    r.rptQuery = rptQuery
    r.Set("rpt_query", rptQuery)
    return nil
}

func (r TaobaoFeedflowAccountRpthourlistRequest) GetRptQuery() *RptQueryDto {
    return r.rptQuery
}

