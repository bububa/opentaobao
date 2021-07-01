package usergrowth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取 手淘用增 线下业务 t+8汇总数据 API请求
taobao.usergrowth.offline.convertion.summary.eight.get

淘系用户增长团队-线下拉新业务，对线下渠道提供mapi，目的是为了给有开发能力的渠道提供返数功能，方便渠道对手下的推广者结算（t+8汇总）
*/
type TaobaoUsergrowthOfflineConvertionSummaryEightGetAPIRequest struct {
    model.Params
    // 入参
    _query   *OfflineMapiQuery
}

// 初始化TaobaoUsergrowthOfflineConvertionSummaryEightGetAPIRequest对象
func NewTaobaoUsergrowthOfflineConvertionSummaryEightGetRequest() *TaobaoUsergrowthOfflineConvertionSummaryEightGetAPIRequest{
    return &TaobaoUsergrowthOfflineConvertionSummaryEightGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoUsergrowthOfflineConvertionSummaryEightGetAPIRequest) GetApiMethodName() string {
    return "taobao.usergrowth.offline.convertion.summary.eight.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoUsergrowthOfflineConvertionSummaryEightGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Query Setter
// 入参
func (r *TaobaoUsergrowthOfflineConvertionSummaryEightGetAPIRequest) SetQuery(_query *OfflineMapiQuery) error {
    r._query = _query
    r.Set("query", _query)
    return nil
}

// Query Getter
func (r TaobaoUsergrowthOfflineConvertionSummaryEightGetAPIRequest) GetQuery() *OfflineMapiQuery {
    return r._query
}
