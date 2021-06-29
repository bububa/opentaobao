package usergrowth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
手淘线下拉新业务 t+8转化明细数据 API请求
taobao.usergrowth.offline.convertion.details.eight.get

手淘线下拉新业务 给合作渠道返回t+8转化明细数据
*/
type TaobaoUsergrowthOfflineConvertionDetailsEightGetRequest struct {
    model.Params
    // 入参
    query   *OfflineMapiQuery
}

// 初始化TaobaoUsergrowthOfflineConvertionDetailsEightGetRequest对象
func NewTaobaoUsergrowthOfflineConvertionDetailsEightGetRequest() *TaobaoUsergrowthOfflineConvertionDetailsEightGetRequest{
    return &TaobaoUsergrowthOfflineConvertionDetailsEightGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoUsergrowthOfflineConvertionDetailsEightGetRequest) GetApiMethodName() string {
    return "taobao.usergrowth.offline.convertion.details.eight.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoUsergrowthOfflineConvertionDetailsEightGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Query Setter
// 入参
func (r *TaobaoUsergrowthOfflineConvertionDetailsEightGetRequest) SetQuery(query *OfflineMapiQuery) error {
    r.query = query
    r.Set("query", query)
    return nil
}

// Query Getter
func (r TaobaoUsergrowthOfflineConvertionDetailsEightGetRequest) GetQuery() *OfflineMapiQuery {
    return r.query
}
