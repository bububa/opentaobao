package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
超级推荐【商品推广】单元分时报表查询 APIRequest
taobao.feedflow.item.adgroup.rpthourlist

广告主推广组分时数据查询，支持广告主查询最近90天内某一天的单元维度分时报表数据
*/
type TaobaoFeedflowItemAdgroupRpthourlistRequest struct {
    model.Params

    // 查询参数
    rptQuery   *RptQueryDto 

}

func NewTaobaoFeedflowItemAdgroupRpthourlistRequest() *TaobaoFeedflowItemAdgroupRpthourlistRequest{
    return &TaobaoFeedflowItemAdgroupRpthourlistRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFeedflowItemAdgroupRpthourlistRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.adgroup.rpthourlist"
}

func (r TaobaoFeedflowItemAdgroupRpthourlistRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFeedflowItemAdgroupRpthourlistRequest) SetRptQuery(rptQuery *RptQueryDto) error {
    r.rptQuery = rptQuery
    r.Set("rpt_query", rptQuery)
    return nil
}

func (r TaobaoFeedflowItemAdgroupRpthourlistRequest) GetRptQuery() *RptQueryDto {
    return r.rptQuery
}

