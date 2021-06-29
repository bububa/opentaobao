package usergrowth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取 手淘用增 线下业务 t+8汇总数据 APIRequest
taobao.usergrowth.offline.convertion.summary.eight.get

淘系用户增长团队-线下拉新业务，对线下渠道提供mapi，目的是为了给有开发能力的渠道提供返数功能，方便渠道对手下的推广者结算（t+8汇总）
*/
type TaobaoUsergrowthOfflineConvertionSummaryEightGetRequest struct {
    model.Params

    // 入参
    query   *OfflineMapiQuery 

}

func NewTaobaoUsergrowthOfflineConvertionSummaryEightGetRequest() *TaobaoUsergrowthOfflineConvertionSummaryEightGetRequest{
    return &TaobaoUsergrowthOfflineConvertionSummaryEightGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoUsergrowthOfflineConvertionSummaryEightGetRequest) GetApiMethodName() string {
    return "taobao.usergrowth.offline.convertion.summary.eight.get"
}

func (r TaobaoUsergrowthOfflineConvertionSummaryEightGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoUsergrowthOfflineConvertionSummaryEightGetRequest) SetQuery(query *OfflineMapiQuery) error {
    r.query = query
    r.Set("query", query)
    return nil
}

func (r TaobaoUsergrowthOfflineConvertionSummaryEightGetRequest) GetQuery() *OfflineMapiQuery {
    return r.query
}

