package auction

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAuctionGovDataMonthlyGetAPIRequest 按月统计法院拍卖数据 API请求
// taobao.auction.gov.data.monthly.get
//
// 按月统计法院拍卖数据
// 包含：
// 标的件数统计：发布标的件数、结束标的件数、开拍标的件数
// 竞价实况：预计成交金额、出价次数、报名人数
// 在线标的：在线标的件数、意向用户数、网拍围观人次
//
// 最长12个月，月的起始时间不能早于2017年3月
type TaobaoAuctionGovDataMonthlyGetAPIRequest struct {
	model.Params
	// 法院名称
	_courtName string
	// 开始月份
	_startMonth string
	// 截止月份(统计数据包含这个月)
	_endMonth string
	// 统计数据是够包含下属法院
	_isIncludeSub bool
}

// NewTaobaoAuctionGovDataMonthlyGetRequest 初始化TaobaoAuctionGovDataMonthlyGetAPIRequest对象
func NewTaobaoAuctionGovDataMonthlyGetRequest() *TaobaoAuctionGovDataMonthlyGetAPIRequest {
	return &TaobaoAuctionGovDataMonthlyGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAuctionGovDataMonthlyGetAPIRequest) GetApiMethodName() string {
	return "taobao.auction.gov.data.monthly.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAuctionGovDataMonthlyGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetCourtName is CourtName Setter
// 法院名称
func (r *TaobaoAuctionGovDataMonthlyGetAPIRequest) SetCourtName(_courtName string) error {
	r._courtName = _courtName
	r.Set("court_name", _courtName)
	return nil
}

// GetCourtName CourtName Getter
func (r TaobaoAuctionGovDataMonthlyGetAPIRequest) GetCourtName() string {
	return r._courtName
}

// SetStartMonth is StartMonth Setter
// 开始月份
func (r *TaobaoAuctionGovDataMonthlyGetAPIRequest) SetStartMonth(_startMonth string) error {
	r._startMonth = _startMonth
	r.Set("start_month", _startMonth)
	return nil
}

// GetStartMonth StartMonth Getter
func (r TaobaoAuctionGovDataMonthlyGetAPIRequest) GetStartMonth() string {
	return r._startMonth
}

// SetEndMonth is EndMonth Setter
// 截止月份(统计数据包含这个月)
func (r *TaobaoAuctionGovDataMonthlyGetAPIRequest) SetEndMonth(_endMonth string) error {
	r._endMonth = _endMonth
	r.Set("end_month", _endMonth)
	return nil
}

// GetEndMonth EndMonth Getter
func (r TaobaoAuctionGovDataMonthlyGetAPIRequest) GetEndMonth() string {
	return r._endMonth
}

// SetIsIncludeSub is IsIncludeSub Setter
// 统计数据是够包含下属法院
func (r *TaobaoAuctionGovDataMonthlyGetAPIRequest) SetIsIncludeSub(_isIncludeSub bool) error {
	r._isIncludeSub = _isIncludeSub
	r.Set("is_include_sub", _isIncludeSub)
	return nil
}

// GetIsIncludeSub IsIncludeSub Getter
func (r TaobaoAuctionGovDataMonthlyGetAPIRequest) GetIsIncludeSub() bool {
	return r._isIncludeSub
}
