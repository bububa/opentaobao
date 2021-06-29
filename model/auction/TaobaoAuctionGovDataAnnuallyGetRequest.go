package auction

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
按年统计法院拍卖数据 API请求
taobao.auction.gov.data.annually.get

按月统计法院拍卖数据 包含：
标的件数统计：发布标的件数、结束标的件数、开拍标的件数
竞价实况：预计成交金额、出价次数、报名人数
在线标的：在线标的件数、意向用户数、网拍围观人次

最长6年，年起始时间2017年
*/
type TaobaoAuctionGovDataAnnuallyGetRequest struct {
    model.Params
    // 法院名称
    courtName   string
    // 统计数据是够包含下属法院
    isIncludeSub   bool
    // 开始年份
    startYear   string
    // 结束年份
    endYear   string
}

// 初始化TaobaoAuctionGovDataAnnuallyGetRequest对象
func NewTaobaoAuctionGovDataAnnuallyGetRequest() *TaobaoAuctionGovDataAnnuallyGetRequest{
    return &TaobaoAuctionGovDataAnnuallyGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAuctionGovDataAnnuallyGetRequest) GetApiMethodName() string {
    return "taobao.auction.gov.data.annually.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAuctionGovDataAnnuallyGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CourtName Setter
// 法院名称
func (r *TaobaoAuctionGovDataAnnuallyGetRequest) SetCourtName(courtName string) error {
    r.courtName = courtName
    r.Set("court_name", courtName)
    return nil
}

// CourtName Getter
func (r TaobaoAuctionGovDataAnnuallyGetRequest) GetCourtName() string {
    return r.courtName
}
// IsIncludeSub Setter
// 统计数据是够包含下属法院
func (r *TaobaoAuctionGovDataAnnuallyGetRequest) SetIsIncludeSub(isIncludeSub bool) error {
    r.isIncludeSub = isIncludeSub
    r.Set("is_include_sub", isIncludeSub)
    return nil
}

// IsIncludeSub Getter
func (r TaobaoAuctionGovDataAnnuallyGetRequest) GetIsIncludeSub() bool {
    return r.isIncludeSub
}
// StartYear Setter
// 开始年份
func (r *TaobaoAuctionGovDataAnnuallyGetRequest) SetStartYear(startYear string) error {
    r.startYear = startYear
    r.Set("start_year", startYear)
    return nil
}

// StartYear Getter
func (r TaobaoAuctionGovDataAnnuallyGetRequest) GetStartYear() string {
    return r.startYear
}
// EndYear Setter
// 结束年份
func (r *TaobaoAuctionGovDataAnnuallyGetRequest) SetEndYear(endYear string) error {
    r.endYear = endYear
    r.Set("end_year", endYear)
    return nil
}

// EndYear Getter
func (r TaobaoAuctionGovDataAnnuallyGetRequest) GetEndYear() string {
    return r.endYear
}
