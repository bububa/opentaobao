package usergrowth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘系用增线下转化明细 API请求
taobao.usergrowth.offline.convertion.details.get

淘系用增增长-线下拉新：为渠道提供返回拉新转化数据接口
*/
type TaobaoUsergrowthOfflineConvertionDetailsGetRequest struct {
    model.Params
    // 入参
    _query   *OfflineMapiQuery
}

// 初始化TaobaoUsergrowthOfflineConvertionDetailsGetRequest对象
func NewTaobaoUsergrowthOfflineConvertionDetailsGetRequest() *TaobaoUsergrowthOfflineConvertionDetailsGetRequest{
    return &TaobaoUsergrowthOfflineConvertionDetailsGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoUsergrowthOfflineConvertionDetailsGetRequest) GetApiMethodName() string {
    return "taobao.usergrowth.offline.convertion.details.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoUsergrowthOfflineConvertionDetailsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Query Setter
// 入参
func (r *TaobaoUsergrowthOfflineConvertionDetailsGetRequest) SetQuery(_query *OfflineMapiQuery) error {
    r._query = _query
    r.Set("query", _query)
    return nil
}

// Query Getter
func (r TaobaoUsergrowthOfflineConvertionDetailsGetRequest) GetQuery() *OfflineMapiQuery {
    return r._query
}
