package crm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCrmMembersGetPrivyAPIRequest 获取卖家会员(基本查询) API请求
// taobao.crm.members.get.privy
//
// 查询卖家的会员，进行基本的查询，返回符合条件的会员列表
type TaobaoCrmMembersGetPrivyAPIRequest struct {
	model.Params
	// 最早上次交易时间
	_minLastTradeTime string
	// 最迟上次交易时间
	_maxLastTradeTime string
	// ouid
	_ouid string
	// 会员等级,如果不传入值则默认为全部等级。
	_grade int64
	// 最小交易额,单位为元
	_minTradeAmount float64
	// 最大交易额，单位为元
	_maxTradeAmount float64
	// 最小交易量
	_minTradeCount int64
	// 最大交易量
	_maxTradeCount int64
	// 表示每页显示的会员数量,page_size的最大值不能超过100条,最小值不能低于1，
	_pageSize int64
	// 显示第几页的会员，如果输入的页码大于总共的页码数，例如总共10页，但是current_page的值为11，则返回空白页，最小页数为1，最大页数为80
	_currentPage int64
}

// NewTaobaoCrmMembersGetPrivyRequest 初始化TaobaoCrmMembersGetPrivyAPIRequest对象
func NewTaobaoCrmMembersGetPrivyRequest() *TaobaoCrmMembersGetPrivyAPIRequest {
	return &TaobaoCrmMembersGetPrivyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoCrmMembersGetPrivyAPIRequest) GetApiMethodName() string {
	return "taobao.crm.members.get.privy"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoCrmMembersGetPrivyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetMinLastTradeTime is MinLastTradeTime Setter
// 最早上次交易时间
func (r *TaobaoCrmMembersGetPrivyAPIRequest) SetMinLastTradeTime(_minLastTradeTime string) error {
	r._minLastTradeTime = _minLastTradeTime
	r.Set("min_last_trade_time", _minLastTradeTime)
	return nil
}

// GetMinLastTradeTime MinLastTradeTime Getter
func (r TaobaoCrmMembersGetPrivyAPIRequest) GetMinLastTradeTime() string {
	return r._minLastTradeTime
}

// SetMaxLastTradeTime is MaxLastTradeTime Setter
// 最迟上次交易时间
func (r *TaobaoCrmMembersGetPrivyAPIRequest) SetMaxLastTradeTime(_maxLastTradeTime string) error {
	r._maxLastTradeTime = _maxLastTradeTime
	r.Set("max_last_trade_time", _maxLastTradeTime)
	return nil
}

// GetMaxLastTradeTime MaxLastTradeTime Getter
func (r TaobaoCrmMembersGetPrivyAPIRequest) GetMaxLastTradeTime() string {
	return r._maxLastTradeTime
}

// SetOuid is Ouid Setter
// ouid
func (r *TaobaoCrmMembersGetPrivyAPIRequest) SetOuid(_ouid string) error {
	r._ouid = _ouid
	r.Set("ouid", _ouid)
	return nil
}

// GetOuid Ouid Getter
func (r TaobaoCrmMembersGetPrivyAPIRequest) GetOuid() string {
	return r._ouid
}

// SetGrade is Grade Setter
// 会员等级,如果不传入值则默认为全部等级。
func (r *TaobaoCrmMembersGetPrivyAPIRequest) SetGrade(_grade int64) error {
	r._grade = _grade
	r.Set("grade", _grade)
	return nil
}

// GetGrade Grade Getter
func (r TaobaoCrmMembersGetPrivyAPIRequest) GetGrade() int64 {
	return r._grade
}

// SetMinTradeAmount is MinTradeAmount Setter
// 最小交易额,单位为元
func (r *TaobaoCrmMembersGetPrivyAPIRequest) SetMinTradeAmount(_minTradeAmount float64) error {
	r._minTradeAmount = _minTradeAmount
	r.Set("min_trade_amount", _minTradeAmount)
	return nil
}

// GetMinTradeAmount MinTradeAmount Getter
func (r TaobaoCrmMembersGetPrivyAPIRequest) GetMinTradeAmount() float64 {
	return r._minTradeAmount
}

// SetMaxTradeAmount is MaxTradeAmount Setter
// 最大交易额，单位为元
func (r *TaobaoCrmMembersGetPrivyAPIRequest) SetMaxTradeAmount(_maxTradeAmount float64) error {
	r._maxTradeAmount = _maxTradeAmount
	r.Set("max_trade_amount", _maxTradeAmount)
	return nil
}

// GetMaxTradeAmount MaxTradeAmount Getter
func (r TaobaoCrmMembersGetPrivyAPIRequest) GetMaxTradeAmount() float64 {
	return r._maxTradeAmount
}

// SetMinTradeCount is MinTradeCount Setter
// 最小交易量
func (r *TaobaoCrmMembersGetPrivyAPIRequest) SetMinTradeCount(_minTradeCount int64) error {
	r._minTradeCount = _minTradeCount
	r.Set("min_trade_count", _minTradeCount)
	return nil
}

// GetMinTradeCount MinTradeCount Getter
func (r TaobaoCrmMembersGetPrivyAPIRequest) GetMinTradeCount() int64 {
	return r._minTradeCount
}

// SetMaxTradeCount is MaxTradeCount Setter
// 最大交易量
func (r *TaobaoCrmMembersGetPrivyAPIRequest) SetMaxTradeCount(_maxTradeCount int64) error {
	r._maxTradeCount = _maxTradeCount
	r.Set("max_trade_count", _maxTradeCount)
	return nil
}

// GetMaxTradeCount MaxTradeCount Getter
func (r TaobaoCrmMembersGetPrivyAPIRequest) GetMaxTradeCount() int64 {
	return r._maxTradeCount
}

// SetPageSize is PageSize Setter
// 表示每页显示的会员数量,page_size的最大值不能超过100条,最小值不能低于1，
func (r *TaobaoCrmMembersGetPrivyAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoCrmMembersGetPrivyAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetCurrentPage is CurrentPage Setter
// 显示第几页的会员，如果输入的页码大于总共的页码数，例如总共10页，但是current_page的值为11，则返回空白页，最小页数为1，最大页数为80
func (r *TaobaoCrmMembersGetPrivyAPIRequest) SetCurrentPage(_currentPage int64) error {
	r._currentPage = _currentPage
	r.Set("current_page", _currentPage)
	return nil
}

// GetCurrentPage CurrentPage Getter
func (r TaobaoCrmMembersGetPrivyAPIRequest) GetCurrentPage() int64 {
	return r._currentPage
}
