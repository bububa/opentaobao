package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
推广计划分日数据查询 API请求
taobao.feedflow.item.campaign.rptdailylist

推广计划分日数据查询
*/
type TaobaoFeedflowItemCampaignRptdailylistAPIRequest struct {
    model.Params
    // 查询条件
    _rptQueryDTO   *RptQueryDto
}

// 初始化TaobaoFeedflowItemCampaignRptdailylistAPIRequest对象
func NewTaobaoFeedflowItemCampaignRptdailylistRequest() *TaobaoFeedflowItemCampaignRptdailylistAPIRequest{
    return &TaobaoFeedflowItemCampaignRptdailylistAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemCampaignRptdailylistAPIRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.campaign.rptdailylist"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemCampaignRptdailylistAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RptQueryDTO Setter
// 查询条件
func (r *TaobaoFeedflowItemCampaignRptdailylistAPIRequest) SetRptQueryDTO(_rptQueryDTO *RptQueryDto) error {
    r._rptQueryDTO = _rptQueryDTO
    r.Set("rpt_query_d_t_o", _rptQueryDTO)
    return nil
}

// RptQueryDTO Getter
func (r TaobaoFeedflowItemCampaignRptdailylistAPIRequest) GetRptQueryDTO() *RptQueryDto {
    return r._rptQueryDTO
}
