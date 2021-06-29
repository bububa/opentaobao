package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
定向分日数据查询 API请求
taobao.feedflow.item.crowd.rptdailylist

定向分日数据查询
*/
type TaobaoFeedflowItemCrowdRptdailylistRequest struct {
    model.Params
    // 查询条件
    rptQueryDTO   *RptQueryDto
}

// 初始化TaobaoFeedflowItemCrowdRptdailylistRequest对象
func NewTaobaoFeedflowItemCrowdRptdailylistRequest() *TaobaoFeedflowItemCrowdRptdailylistRequest{
    return &TaobaoFeedflowItemCrowdRptdailylistRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemCrowdRptdailylistRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.crowd.rptdailylist"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemCrowdRptdailylistRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RptQueryDTO Setter
// 查询条件
func (r *TaobaoFeedflowItemCrowdRptdailylistRequest) SetRptQueryDTO(rptQueryDTO *RptQueryDto) error {
    r.rptQueryDTO = rptQueryDTO
    r.Set("rpt_query_d_t_o", rptQueryDTO)
    return nil
}

// RptQueryDTO Getter
func (r TaobaoFeedflowItemCrowdRptdailylistRequest) GetRptQueryDTO() *RptQueryDto {
    return r.rptQueryDTO
}
