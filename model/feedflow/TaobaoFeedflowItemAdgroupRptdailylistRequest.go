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
type TaobaoFeedflowItemAdgroupRptdailylistRequest struct {
    model.Params
    // 查询条件
    _rptQueryDTO   *RptQueryDto
}

// 初始化TaobaoFeedflowItemAdgroupRptdailylistRequest对象
func NewTaobaoFeedflowItemAdgroupRptdailylistRequest() *TaobaoFeedflowItemAdgroupRptdailylistRequest{
    return &TaobaoFeedflowItemAdgroupRptdailylistRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemAdgroupRptdailylistRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.adgroup.rptdailylist"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemAdgroupRptdailylistRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RptQueryDTO Setter
// 查询条件
func (r *TaobaoFeedflowItemAdgroupRptdailylistRequest) SetRptQueryDTO(_rptQueryDTO *RptQueryDto) error {
    r._rptQueryDTO = _rptQueryDTO
    r.Set("rpt_query_d_t_o", _rptQueryDTO)
    return nil
}

// RptQueryDTO Getter
func (r TaobaoFeedflowItemAdgroupRptdailylistRequest) GetRptQueryDTO() *RptQueryDto {
    return r._rptQueryDTO
}
