package usergrowth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
手淘线下拉新业务 t+8转化明细数据 APIRequest
taobao.usergrowth.offline.convertion.details.eight.get

手淘线下拉新业务 给合作渠道返回t+8转化明细数据
*/
type TaobaoUsergrowthOfflineConvertionDetailsEightGetRequest struct {
    model.Params

    // 入参
    query   *OfflineMapiQuery 

}

func NewTaobaoUsergrowthOfflineConvertionDetailsEightGetRequest() *TaobaoUsergrowthOfflineConvertionDetailsEightGetRequest{
    return &TaobaoUsergrowthOfflineConvertionDetailsEightGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoUsergrowthOfflineConvertionDetailsEightGetRequest) GetApiMethodName() string {
    return "taobao.usergrowth.offline.convertion.details.eight.get"
}

func (r TaobaoUsergrowthOfflineConvertionDetailsEightGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoUsergrowthOfflineConvertionDetailsEightGetRequest) SetQuery(query *OfflineMapiQuery) error {
    r.query = query
    r.Set("query", query)
    return nil
}

func (r TaobaoUsergrowthOfflineConvertionDetailsEightGetRequest) GetQuery() *OfflineMapiQuery {
    return r.query
}

