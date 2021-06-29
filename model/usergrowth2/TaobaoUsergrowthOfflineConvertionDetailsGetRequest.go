package usergrowth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘系用增线下转化明细 APIRequest
taobao.usergrowth.offline.convertion.details.get

淘系用增增长-线下拉新：为渠道提供返回拉新转化数据接口
*/
type TaobaoUsergrowthOfflineConvertionDetailsGetRequest struct {
    model.Params

    // 入参
    query   *OfflineMapiQuery 

}

func NewTaobaoUsergrowthOfflineConvertionDetailsGetRequest() *TaobaoUsergrowthOfflineConvertionDetailsGetRequest{
    return &TaobaoUsergrowthOfflineConvertionDetailsGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoUsergrowthOfflineConvertionDetailsGetRequest) GetApiMethodName() string {
    return "taobao.usergrowth.offline.convertion.details.get"
}

func (r TaobaoUsergrowthOfflineConvertionDetailsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoUsergrowthOfflineConvertionDetailsGetRequest) SetQuery(query *OfflineMapiQuery) error {
    r.query = query
    r.Set("query", query)
    return nil
}

func (r TaobaoUsergrowthOfflineConvertionDetailsGetRequest) GetQuery() *OfflineMapiQuery {
    return r.query
}

