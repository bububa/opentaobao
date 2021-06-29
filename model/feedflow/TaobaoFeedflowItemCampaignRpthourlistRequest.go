package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
超级推荐【商品推广】计划分时报表查询 APIRequest
taobao.feedflow.item.campaign.rpthourlist

广告主推广计划分时数据查询，支持广告主查询最近90天内某一天的计划维度分时报表数据
*/
type TaobaoFeedflowItemCampaignRpthourlistRequest struct {
    model.Params

    // 查询参数
    rptQuery   *RptQueryDto 

}

func NewTaobaoFeedflowItemCampaignRpthourlistRequest() *TaobaoFeedflowItemCampaignRpthourlistRequest{
    return &TaobaoFeedflowItemCampaignRpthourlistRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFeedflowItemCampaignRpthourlistRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.campaign.rpthourlist"
}

func (r TaobaoFeedflowItemCampaignRpthourlistRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFeedflowItemCampaignRpthourlistRequest) SetRptQuery(rptQuery *RptQueryDto) error {
    r.rptQuery = rptQuery
    r.Set("rpt_query", rptQuery)
    return nil
}

func (r TaobaoFeedflowItemCampaignRpthourlistRequest) GetRptQuery() *RptQueryDto {
    return r.rptQuery
}

