package auction

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取实时(今日)统计数据 API请求
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

// 初始化TaobaoAuctionGovDataRealtimeGetRequest对象
func NewTaobaoAuctionGovDataRealtimeGetRequest() *TaobaoAuctionGovDataRealtimeGetRequest{
    return &TaobaoAuctionGovDataRealtimeGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAuctionGovDataRealtimeGetRequest) GetApiMethodName() string {
    return "taobao.auction.gov.data.realtime.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAuctionGovDataRealtimeGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CourtName Setter
// 法院名称
func (r *TaobaoAuctionGovDataRealtimeGetRequest) SetCourtName(courtName string) error {
    r.courtName = courtName
    r.Set("court_name", courtName)
    return nil
}

// CourtName Getter
func (r TaobaoAuctionGovDataRealtimeGetRequest) GetCourtName() string {
    return r.courtName
}
// IsIncludeSub Setter
// 统计数据是否包含下级法院
func (r *TaobaoAuctionGovDataRealtimeGetRequest) SetIsIncludeSub(isIncludeSub bool) error {
    r.isIncludeSub = isIncludeSub
    r.Set("is_include_sub", isIncludeSub)
    return nil
}

// IsIncludeSub Getter
func (r TaobaoAuctionGovDataRealtimeGetRequest) GetIsIncludeSub() bool {
    return r.isIncludeSub
}
