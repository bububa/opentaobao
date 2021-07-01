package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
超级推荐【商品推广】单元分时报表查询 API请求
taobao.feedflow.item.adgroup.rpthourlist

广告主推广组分时数据查询，支持广告主查询最近90天内某一天的单元维度分时报表数据
*/
type TaobaoFeedflowItemAdgroupRpthourlistAPIRequest struct {
    model.Params
    // 查询参数
    _rptQuery   *RptQueryDTO
}

// 初始化TaobaoFeedflowItemAdgroupRpthourlistAPIRequest对象
func NewTaobaoFeedflowItemAdgroupRpthourlistRequest() *TaobaoFeedflowItemAdgroupRpthourlistAPIRequest{
    return &TaobaoFeedflowItemAdgroupRpthourlistAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemAdgroupRpthourlistAPIRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.adgroup.rpthourlist"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemAdgroupRpthourlistAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RptQuery Setter
// 查询参数
func (r *TaobaoFeedflowItemAdgroupRpthourlistAPIRequest) SetRptQuery(_rptQuery *RptQueryDTO) error {
    r._rptQuery = _rptQuery
    r.Set("rpt_query", _rptQuery)
    return nil
}

// RptQuery Getter
func (r TaobaoFeedflowItemAdgroupRpthourlistAPIRequest) GetRptQuery() *RptQueryDTO {
    return r._rptQuery
}
