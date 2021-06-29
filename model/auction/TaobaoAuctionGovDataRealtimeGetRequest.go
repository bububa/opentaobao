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
    _courtName   string
    // 统计数据是否包含下级法院
    _isIncludeSub   bool
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
func (r *TaobaoAuctionGovDataRealtimeGetRequest) SetCourtName(_courtName string) error {
    r._courtName = _courtName
    r.Set("court_name", _courtName)
    return nil
}

// CourtName Getter
func (r TaobaoAuctionGovDataRealtimeGetRequest) GetCourtName() string {
    return r._courtName
}
// IsIncludeSub Setter
// 统计数据是否包含下级法院
func (r *TaobaoAuctionGovDataRealtimeGetRequest) SetIsIncludeSub(_isIncludeSub bool) error {
    r._isIncludeSub = _isIncludeSub
    r.Set("is_include_sub", _isIncludeSub)
    return nil
}

// IsIncludeSub Getter
func (r TaobaoAuctionGovDataRealtimeGetRequest) GetIsIncludeSub() bool {
    return r._isIncludeSub
}
