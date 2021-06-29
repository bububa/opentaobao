package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
推广计划分日数据查询 APIRequest
taobao.feedflow.item.campaign.rptdailylist

推广计划分日数据查询
*/
type TaobaoFeedflowItemCampaignRptdailylistRequest struct {
    model.Params

    // 查询条件
    rptQueryDTO   *RptQueryDto 

}

func NewTaobaoFeedflowItemCampaignRptdailylistRequest() *TaobaoFeedflowItemCampaignRptdailylistRequest{
    return &TaobaoFeedflowItemCampaignRptdailylistRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFeedflowItemCampaignRptdailylistRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.campaign.rptdailylist"
}

func (r TaobaoFeedflowItemCampaignRptdailylistRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFeedflowItemCampaignRptdailylistRequest) SetRptQueryDTO(rptQueryDTO *RptQueryDto) error {
    r.rptQueryDTO = rptQueryDTO
    r.Set("rpt_query_d_t_o", rptQueryDTO)
    return nil
}

func (r TaobaoFeedflowItemCampaignRptdailylistRequest) GetRptQueryDTO() *RptQueryDto {
    return r.rptQueryDTO
}

