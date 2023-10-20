package auction

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoauctiongovdataannuallygetAPIRequest 按年统计法院拍卖数据 API请求
// taobao.auction.gov.data.annually.get
//
// 按月统计法院拍卖数据 包含：
// 标的件数统计：发布标的件数、结束标的件数、开拍标的件数
// 竞价实况：预计成交金额、出价次数、报名人数
// 在线标的：在线标的件数、意向用户数、网拍围观人次
//
// 最长6年，年起始时间2017年
type TaobaoauctiongovdataannuallygetAPIRequest struct {
	model.Params
	// 法院名称
	_courtName string
	// 开始年份
	_startYear string
	// 结束年份
	_endYear string
	// 统计数据是够包含下属法院
	_isIncludeSub bool
}

// NewTaobaoauctiongovdataannuallygetRequest 初始化TaobaoauctiongovdataannuallygetAPIRequest对象
func NewTaobaoauctiongovdataannuallygetRequest() *TaobaoauctiongovdataannuallygetAPIRequest {
	return &TaobaoauctiongovdataannuallygetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoauctiongovdataannuallygetAPIRequest) GetApiMethodName() string {
	return "taobao.auction.gov.data.annually.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoauctiongovdataannuallygetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoauctiongovdataannuallygetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCourtName is CourtName Setter
// 法院名称
func (r *TaobaoauctiongovdataannuallygetAPIRequest) SetCourtName(_courtName string) error {
	r._courtName = _courtName
	r.Set("court_name", _courtName)
	return nil
}

// GetCourtName CourtName Getter
func (r TaobaoauctiongovdataannuallygetAPIRequest) GetCourtName() string {
	return r._courtName
}

// SetStartYear is StartYear Setter
// 开始年份
func (r *TaobaoauctiongovdataannuallygetAPIRequest) SetStartYear(_startYear string) error {
	r._startYear = _startYear
	r.Set("start_year", _startYear)
	return nil
}

// GetStartYear StartYear Getter
func (r TaobaoauctiongovdataannuallygetAPIRequest) GetStartYear() string {
	return r._startYear
}

// SetEndYear is EndYear Setter
// 结束年份
func (r *TaobaoauctiongovdataannuallygetAPIRequest) SetEndYear(_endYear string) error {
	r._endYear = _endYear
	r.Set("end_year", _endYear)
	return nil
}

// GetEndYear EndYear Getter
func (r TaobaoauctiongovdataannuallygetAPIRequest) GetEndYear() string {
	return r._endYear
}

// SetIsIncludeSub is IsIncludeSub Setter
// 统计数据是够包含下属法院
func (r *TaobaoauctiongovdataannuallygetAPIRequest) SetIsIncludeSub(_isIncludeSub bool) error {
	r._isIncludeSub = _isIncludeSub
	r.Set("is_include_sub", _isIncludeSub)
	return nil
}

// GetIsIncludeSub IsIncludeSub Getter
func (r TaobaoauctiongovdataannuallygetAPIRequest) GetIsIncludeSub() bool {
	return r._isIncludeSub
}
