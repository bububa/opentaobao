package auction

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据不同维度，获取排行榜列表 API请求
taobao.auction.gov.data.topn.get

根据不同时间维度(周,月,年)，获取(成交额或发拍件数)排行榜列表
*/
type TaobaoAuctionGovDataTopnGetRequest struct {
    model.Params
    // 周期类型  （2：周，3：月，4：年）
    circleType   int64
    // 周期区间 周期（周填0、月份 yyyyMM、年份 yyyy)
    circle   int64
    // 业务类型 （1：成交额，2：发拍件数）
    busiType   int64
    // 区域类型（1：全国，2：全省）
    zoneType   int64
    // 法院名称
    courtName   string
}

// 初始化TaobaoAuctionGovDataTopnGetRequest对象
func NewTaobaoAuctionGovDataTopnGetRequest() *TaobaoAuctionGovDataTopnGetRequest{
    return &TaobaoAuctionGovDataTopnGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAuctionGovDataTopnGetRequest) GetApiMethodName() string {
    return "taobao.auction.gov.data.topn.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAuctionGovDataTopnGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CircleType Setter
// 周期类型  （2：周，3：月，4：年）
func (r *TaobaoAuctionGovDataTopnGetRequest) SetCircleType(circleType int64) error {
    r.circleType = circleType
    r.Set("circle_type", circleType)
    return nil
}

// CircleType Getter
func (r TaobaoAuctionGovDataTopnGetRequest) GetCircleType() int64 {
    return r.circleType
}
// Circle Setter
// 周期区间 周期（周填0、月份 yyyyMM、年份 yyyy)
func (r *TaobaoAuctionGovDataTopnGetRequest) SetCircle(circle int64) error {
    r.circle = circle
    r.Set("circle", circle)
    return nil
}

// Circle Getter
func (r TaobaoAuctionGovDataTopnGetRequest) GetCircle() int64 {
    return r.circle
}
// BusiType Setter
// 业务类型 （1：成交额，2：发拍件数）
func (r *TaobaoAuctionGovDataTopnGetRequest) SetBusiType(busiType int64) error {
    r.busiType = busiType
    r.Set("busi_type", busiType)
    return nil
}

// BusiType Getter
func (r TaobaoAuctionGovDataTopnGetRequest) GetBusiType() int64 {
    return r.busiType
}
// ZoneType Setter
// 区域类型（1：全国，2：全省）
func (r *TaobaoAuctionGovDataTopnGetRequest) SetZoneType(zoneType int64) error {
    r.zoneType = zoneType
    r.Set("zone_type", zoneType)
    return nil
}

// ZoneType Getter
func (r TaobaoAuctionGovDataTopnGetRequest) GetZoneType() int64 {
    return r.zoneType
}
// CourtName Setter
// 法院名称
func (r *TaobaoAuctionGovDataTopnGetRequest) SetCourtName(courtName string) error {
    r.courtName = courtName
    r.Set("court_name", courtName)
    return nil
}

// CourtName Getter
func (r TaobaoAuctionGovDataTopnGetRequest) GetCourtName() string {
    return r.courtName
}
