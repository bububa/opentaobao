package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创意分日数据查询 API请求
taobao.feedflow.item.creative.rptdailylist

创意分日数据查询
*/
type TaobaoFeedflowItemCreativeRptdailylistRequest struct {
    model.Params
    // 查询条件
    rptQueryDTO   *RptQueryDto
}

// 初始化TaobaoFeedflowItemCreativeRptdailylistRequest对象
func NewTaobaoFeedflowItemCreativeRptdailylistRequest() *TaobaoFeedflowItemCreativeRptdailylistRequest{
    return &TaobaoFeedflowItemCreativeRptdailylistRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemCreativeRptdailylistRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.creative.rptdailylist"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemCreativeRptdailylistRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RptQueryDTO Setter
// 查询条件
func (r *TaobaoFeedflowItemCreativeRptdailylistRequest) SetRptQueryDTO(rptQueryDTO *RptQueryDto) error {
    r.rptQueryDTO = rptQueryDTO
    r.Set("rpt_query_d_t_o", rptQueryDTO)
    return nil
}

// RptQueryDTO Getter
func (r TaobaoFeedflowItemCreativeRptdailylistRequest) GetRptQueryDTO() *RptQueryDto {
    return r.rptQueryDTO
}
