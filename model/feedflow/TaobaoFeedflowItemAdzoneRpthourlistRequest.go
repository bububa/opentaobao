package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
超级推荐【商品推广】资源位分时报表查询 API请求
taobao.feedflow.item.adzone.rpthourlist

广告主资源包分时数据查询，支持广告主查询最近90天内某一天的资源包维度分时报表数据
*/
type TaobaoFeedflowItemAdzoneRpthourlistRequest struct {
    model.Params
    // 查询参数
    _rptQuery   *RptQueryDTO
}

// 初始化TaobaoFeedflowItemAdzoneRpthourlistRequest对象
func NewTaobaoFeedflowItemAdzoneRpthourlistRequest() *TaobaoFeedflowItemAdzoneRpthourlistRequest{
    return &TaobaoFeedflowItemAdzoneRpthourlistRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemAdzoneRpthourlistRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.adzone.rpthourlist"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemAdzoneRpthourlistRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RptQuery Setter
// 查询参数
func (r *TaobaoFeedflowItemAdzoneRpthourlistRequest) SetRptQuery(_rptQuery *RptQueryDTO) error {
    r._rptQuery = _rptQuery
    r.Set("rpt_query", _rptQuery)
    return nil
}

// RptQuery Getter
func (r TaobaoFeedflowItemAdzoneRpthourlistRequest) GetRptQuery() *RptQueryDTO {
    return r._rptQuery
}
