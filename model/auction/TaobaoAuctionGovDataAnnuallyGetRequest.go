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
    _courtName   string
    // 统计数据是够包含下属法院
    _isIncludeSub   bool
    // 开始年份
    _startYear   string
    // 结束年份
    _endYear   string
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
func (r *TaobaoAuctionGovDataAnnuallyGetRequest) SetCourtName(_courtName string) error {
    r._courtName = _courtName
    r.Set("court_name", _courtName)
    return nil
}

// CourtName Getter
func (r TaobaoAuctionGovDataAnnuallyGetRequest) GetCourtName() string {
    return r._courtName
}
// IsIncludeSub Setter
// 统计数据是够包含下属法院
func (r *TaobaoAuctionGovDataAnnuallyGetRequest) SetIsIncludeSub(_isIncludeSub bool) error {
    r._isIncludeSub = _isIncludeSub
    r.Set("is_include_sub", _isIncludeSub)
    return nil
}

// IsIncludeSub Getter
func (r TaobaoAuctionGovDataAnnuallyGetRequest) GetIsIncludeSub() bool {
    return r._isIncludeSub
}
// StartYear Setter
// 开始年份
func (r *TaobaoAuctionGovDataAnnuallyGetRequest) SetStartYear(_startYear string) error {
    r._startYear = _startYear
    r.Set("start_year", _startYear)
    return nil
}

// StartYear Getter
func (r TaobaoAuctionGovDataAnnuallyGetRequest) GetStartYear() string {
    return r._startYear
}
// EndYear Setter
// 结束年份
func (r *TaobaoAuctionGovDataAnnuallyGetRequest) SetEndYear(_endYear string) error {
    r._endYear = _endYear
    r.Set("end_year", _endYear)
    return nil
}

// EndYear Getter
func (r TaobaoAuctionGovDataAnnuallyGetRequest) GetEndYear() string {
    return r._endYear
}
