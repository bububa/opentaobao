package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
推广单元分日数据查询 API请求
taobao.feedflow.item.adgroup.rptdailylist

推广单元分日数据查询
*/
type TaobaoFeedflowItemAdgroupRptdailylistAPIRequest struct {
    model.Params
    // 查询条件
    _rptQueryDTO   *RptQueryDTO
}

// 初始化TaobaoFeedflowItemAdgroupRptdailylistAPIRequest对象
func NewTaobaoFeedflowItemAdgroupRptdailylistRequest() *TaobaoFeedflowItemAdgroupRptdailylistAPIRequest{
    return &TaobaoFeedflowItemAdgroupRptdailylistAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemAdgroupRptdailylistAPIRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.adgroup.rptdailylist"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemAdgroupRptdailylistAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RptQueryDTO Setter
// 查询条件
func (r *TaobaoFeedflowItemAdgroupRptdailylistAPIRequest) SetRptQueryDTO(_rptQueryDTO *RptQueryDTO) error {
    r._rptQueryDTO = _rptQueryDTO
    r.Set("rpt_query_d_t_o", _rptQueryDTO)
    return nil
}

// RptQueryDTO Getter
func (r TaobaoFeedflowItemAdgroupRptdailylistAPIRequest) GetRptQueryDTO() *RptQueryDTO {
    return r._rptQueryDTO
}
