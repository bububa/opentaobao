package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
超级推荐【商品推广】创意分时报表查询 API请求
taobao.feedflow.item.creative.rpthourlist

创意分时数据查询，支持广告主查询最近90天内某一天的创意维度分时报表数据
*/
type TaobaoFeedflowItemCreativeRpthourlistAPIRequest struct {
    model.Params
    // 查询参数
    _rptQuery   *RptQueryDTO
}

// 初始化TaobaoFeedflowItemCreativeRpthourlistAPIRequest对象
func NewTaobaoFeedflowItemCreativeRpthourlistRequest() *TaobaoFeedflowItemCreativeRpthourlistAPIRequest{
    return &TaobaoFeedflowItemCreativeRpthourlistAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemCreativeRpthourlistAPIRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.creative.rpthourlist"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemCreativeRpthourlistAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RptQuery Setter
// 查询参数
func (r *TaobaoFeedflowItemCreativeRpthourlistAPIRequest) SetRptQuery(_rptQuery *RptQueryDTO) error {
    r._rptQuery = _rptQuery
    r.Set("rpt_query", _rptQuery)
    return nil
}

// RptQuery Getter
func (r TaobaoFeedflowItemCreativeRpthourlistAPIRequest) GetRptQuery() *RptQueryDTO {
    return r._rptQuery
}
