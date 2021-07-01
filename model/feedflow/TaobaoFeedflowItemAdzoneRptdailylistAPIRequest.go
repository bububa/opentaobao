package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
资源包分日数据查询 API请求
taobao.feedflow.item.adzone.rptdailylist

资源包分日数据查询
*/
type TaobaoFeedflowItemAdzoneRptdailylistAPIRequest struct {
    model.Params
    // 查询参数
    _rptQueryDTO   *RptQueryDto
}

// 初始化TaobaoFeedflowItemAdzoneRptdailylistAPIRequest对象
func NewTaobaoFeedflowItemAdzoneRptdailylistRequest() *TaobaoFeedflowItemAdzoneRptdailylistAPIRequest{
    return &TaobaoFeedflowItemAdzoneRptdailylistAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemAdzoneRptdailylistAPIRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.adzone.rptdailylist"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemAdzoneRptdailylistAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RptQueryDTO Setter
// 查询参数
func (r *TaobaoFeedflowItemAdzoneRptdailylistAPIRequest) SetRptQueryDTO(_rptQueryDTO *RptQueryDto) error {
    r._rptQueryDTO = _rptQueryDTO
    r.Set("rpt_query_d_t_o", _rptQueryDTO)
    return nil
}

// RptQueryDTO Getter
func (r TaobaoFeedflowItemAdzoneRptdailylistAPIRequest) GetRptQueryDTO() *RptQueryDto {
    return r._rptQueryDTO
}
