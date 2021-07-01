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
type TaobaoFeedflowItemCreativeRptdailylistAPIRequest struct {
    model.Params
    // 查询条件
    _rptQueryDTO   *RptQueryDTO
}

// 初始化TaobaoFeedflowItemCreativeRptdailylistAPIRequest对象
func NewTaobaoFeedflowItemCreativeRptdailylistRequest() *TaobaoFeedflowItemCreativeRptdailylistAPIRequest{
    return &TaobaoFeedflowItemCreativeRptdailylistAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemCreativeRptdailylistAPIRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.creative.rptdailylist"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemCreativeRptdailylistAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RptQueryDTO Setter
// 查询条件
func (r *TaobaoFeedflowItemCreativeRptdailylistAPIRequest) SetRptQueryDTO(_rptQueryDTO *RptQueryDTO) error {
    r._rptQueryDTO = _rptQueryDTO
    r.Set("rpt_query_d_t_o", _rptQueryDTO)
    return nil
}

// RptQueryDTO Getter
func (r TaobaoFeedflowItemCreativeRptdailylistAPIRequest) GetRptQueryDTO() *RptQueryDTO {
    return r._rptQueryDTO
}
