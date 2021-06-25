package auction

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取实时(今日)统计数据 APIRequest
taobao.auction.gov.data.realtime.get

提供查询当日法院及下属法院的拍卖统计数据
*/
type TaobaoAuctionGovDataRealtimeGetRequest struct {
    model.Params

    // 法院名称
    courtName   string 

    // 统计数据是否包含下级法院
    isIncludeSub   bool 

}

func NewTaobaoAuctionGovDataRealtimeGetRequest() *TaobaoAuctionGovDataRealtimeGetRequest{
    return &TaobaoAuctionGovDataRealtimeGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAuctionGovDataRealtimeGetRequest) GetApiMethodName() string {
    return "taobao.auction.gov.data.realtime.get"
}

func (r TaobaoAuctionGovDataRealtimeGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAuctionGovDataRealtimeGetRequest) SetCourtName(courtName string) error {
    r.courtName = courtName
    r.Set("court_name", courtName)
    return nil
}

func (r TaobaoAuctionGovDataRealtimeGetRequest) GetCourtName() string {
    return r.courtName
}

func (r *TaobaoAuctionGovDataRealtimeGetRequest) SetIsIncludeSub(isIncludeSub bool) error {
    r.isIncludeSub = isIncludeSub
    r.Set("is_include_sub", isIncludeSub)
    return nil
}

func (r TaobaoAuctionGovDataRealtimeGetRequest) GetIsIncludeSub() bool {
    return r.isIncludeSub
}

