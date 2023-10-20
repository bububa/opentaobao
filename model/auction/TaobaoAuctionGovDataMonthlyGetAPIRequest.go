package auction

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoauctiongovdatamonthlygetAPIRequest 按月统计法院拍卖数据 API请求
// taobao.auction.gov.data.monthly.get
//
// 按月统计法院拍卖数据
// 包含：
// 标的件数统计：发布标的件数、结束标的件数、开拍标的件数
// 竞价实况：预计成交金额、出价次数、报名人数
// 在线标的：在线标的件数、意向用户数、网拍围观人次
//
// 最长12个月，月的起始时间不能早于2017年3月
type TaobaoauctiongovdatamonthlygetAPIRequest struct {
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

// NewTaobaoauctiongovdatamonthlygetRequest 初始化TaobaoauctiongovdatamonthlygetAPIRequest对象
func NewTaobaoauctiongovdatamonthlygetRequest() *TaobaoauctiongovdatamonthlygetAPIRequest {
	return &TaobaoauctiongovdatamonthlygetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoauctiongovdatamonthlygetAPIRequest) GetApiMethodName() string {
	return "taobao.auction.gov.data.monthly.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoauctiongovdatamonthlygetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoauctiongovdatamonthlygetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCourtName is CourtName Setter
// 法院名称
func (r *TaobaoauctiongovdatamonthlygetAPIRequest) SetCourtName(_courtName string) error {
	r._courtName = _courtName
	r.Set("court_name", _courtName)
	return nil
}

// GetCourtName CourtName Getter
func (r TaobaoauctiongovdatamonthlygetAPIRequest) GetCourtName() string {
	return r._courtName
}

// SetStartMonth is StartMonth Setter
// 开始月份
func (r *TaobaoauctiongovdatamonthlygetAPIRequest) SetStartMonth(_startMonth string) error {
	r._startMonth = _startMonth
	r.Set("start_month", _startMonth)
	return nil
}

// GetStartMonth StartMonth Getter
func (r TaobaoauctiongovdatamonthlygetAPIRequest) GetStartMonth() string {
	return r._startMonth
}

// SetEndMonth is EndMonth Setter
// 截止月份(统计数据包含这个月)
func (r *TaobaoauctiongovdatamonthlygetAPIRequest) SetEndMonth(_endMonth string) error {
	r._endMonth = _endMonth
	r.Set("end_month", _endMonth)
	return nil
}

// GetEndMonth EndMonth Getter
func (r TaobaoauctiongovdatamonthlygetAPIRequest) GetEndMonth() string {
	return r._endMonth
}

// SetIsIncludeSub is IsIncludeSub Setter
// 统计数据是够包含下属法院
func (r *TaobaoauctiongovdatamonthlygetAPIRequest) SetIsIncludeSub(_isIncludeSub bool) error {
	r._isIncludeSub = _isIncludeSub
	r.Set("is_include_sub", _isIncludeSub)
	return nil
}

// GetIsIncludeSub IsIncludeSub Getter
func (r TaobaoauctiongovdatamonthlygetAPIRequest) GetIsIncludeSub() bool {
	return r._isIncludeSub
}
