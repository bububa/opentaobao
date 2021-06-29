package auction

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
按月统计法院拍卖数据 API请求
taobao.auction.gov.data.monthly.get

按月统计法院拍卖数据 
包含：
标的件数统计：发布标的件数、结束标的件数、开拍标的件数
竞价实况：预计成交金额、出价次数、报名人数
在线标的：在线标的件数、意向用户数、网拍围观人次

最长12个月，月的起始时间不能早于2017年3月
*/
type TaobaoAuctionGovDataMonthlyGetRequest struct {
    model.Params
    // 法院名称
    courtName   string
    // 统计数据是够包含下属法院
    isIncludeSub   bool
    // 开始月份
    startMonth   string
    // 截止月份(统计数据包含这个月)
    endMonth   string
}

// 初始化TaobaoAuctionGovDataMonthlyGetRequest对象
func NewTaobaoAuctionGovDataMonthlyGetRequest() *TaobaoAuctionGovDataMonthlyGetRequest{
    return &TaobaoAuctionGovDataMonthlyGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAuctionGovDataMonthlyGetRequest) GetApiMethodName() string {
    return "taobao.auction.gov.data.monthly.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAuctionGovDataMonthlyGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CourtName Setter
// 法院名称
func (r *TaobaoAuctionGovDataMonthlyGetRequest) SetCourtName(courtName string) error {
    r.courtName = courtName
    r.Set("court_name", courtName)
    return nil
}

// CourtName Getter
func (r TaobaoAuctionGovDataMonthlyGetRequest) GetCourtName() string {
    return r.courtName
}
// IsIncludeSub Setter
// 统计数据是够包含下属法院
func (r *TaobaoAuctionGovDataMonthlyGetRequest) SetIsIncludeSub(isIncludeSub bool) error {
    r.isIncludeSub = isIncludeSub
    r.Set("is_include_sub", isIncludeSub)
    return nil
}

// IsIncludeSub Getter
func (r TaobaoAuctionGovDataMonthlyGetRequest) GetIsIncludeSub() bool {
    return r.isIncludeSub
}
// StartMonth Setter
// 开始月份
func (r *TaobaoAuctionGovDataMonthlyGetRequest) SetStartMonth(startMonth string) error {
    r.startMonth = startMonth
    r.Set("start_month", startMonth)
    return nil
}

// StartMonth Getter
func (r TaobaoAuctionGovDataMonthlyGetRequest) GetStartMonth() string {
    return r.startMonth
}
// EndMonth Setter
// 截止月份(统计数据包含这个月)
func (r *TaobaoAuctionGovDataMonthlyGetRequest) SetEndMonth(endMonth string) error {
    r.endMonth = endMonth
    r.Set("end_month", endMonth)
    return nil
}

// EndMonth Getter
func (r TaobaoAuctionGovDataMonthlyGetRequest) GetEndMonth() string {
    return r.endMonth
}
