package crm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取卖家的会员（基本查询） API请求
taobao.crm.members.get

查询卖家的会员，进行基本的查询，返回符合条件的会员列表
*/
type TaobaoCrmMembersGetRequest struct {
    model.Params
    // 买家的昵称
    _buyerNick   string
    // 会员等级,如果不传入值则默认为全部等级。
    _grade   int64
    // 最小交易额,单位为元
    _minTradeAmount   float64
    // 最大交易额，单位为元
    _maxTradeAmount   float64
    // 最小交易量
    _minTradeCount   int64
    // 最大交易量
    _maxTradeCount   int64
    // 最早上次交易时间
    _minLastTradeTime   string
    // 最迟上次交易时间
    _maxLastTradeTime   string
    // 表示每页显示的会员数量,page_size的最大值不能超过100条,最小值不能低于1，
    _pageSize   int64
    // 显示第几页的会员，如果输入的页码大于总共的页码数，例如总共10页，但是current_page的值为11，则返回空白页，最小页数为1，最大页数为1000
    _currentPage   int64
}

// 初始化TaobaoCrmMembersGetRequest对象
func NewTaobaoCrmMembersGetRequest() *TaobaoCrmMembersGetRequest{
    return &TaobaoCrmMembersGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoCrmMembersGetRequest) GetApiMethodName() string {
    return "taobao.crm.members.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoCrmMembersGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BuyerNick Setter
// 买家的昵称
func (r *TaobaoCrmMembersGetRequest) SetBuyerNick(_buyerNick string) error {
    r._buyerNick = _buyerNick
    r.Set("buyer_nick", _buyerNick)
    return nil
}

// BuyerNick Getter
func (r TaobaoCrmMembersGetRequest) GetBuyerNick() string {
    return r._buyerNick
}
// Grade Setter
// 会员等级,如果不传入值则默认为全部等级。
func (r *TaobaoCrmMembersGetRequest) SetGrade(_grade int64) error {
    r._grade = _grade
    r.Set("grade", _grade)
    return nil
}

// Grade Getter
func (r TaobaoCrmMembersGetRequest) GetGrade() int64 {
    return r._grade
}
// MinTradeAmount Setter
// 最小交易额,单位为元
func (r *TaobaoCrmMembersGetRequest) SetMinTradeAmount(_minTradeAmount float64) error {
    r._minTradeAmount = _minTradeAmount
    r.Set("min_trade_amount", _minTradeAmount)
    return nil
}

// MinTradeAmount Getter
func (r TaobaoCrmMembersGetRequest) GetMinTradeAmount() float64 {
    return r._minTradeAmount
}
// MaxTradeAmount Setter
// 最大交易额，单位为元
func (r *TaobaoCrmMembersGetRequest) SetMaxTradeAmount(_maxTradeAmount float64) error {
    r._maxTradeAmount = _maxTradeAmount
    r.Set("max_trade_amount", _maxTradeAmount)
    return nil
}

// MaxTradeAmount Getter
func (r TaobaoCrmMembersGetRequest) GetMaxTradeAmount() float64 {
    return r._maxTradeAmount
}
// MinTradeCount Setter
// 最小交易量
func (r *TaobaoCrmMembersGetRequest) SetMinTradeCount(_minTradeCount int64) error {
    r._minTradeCount = _minTradeCount
    r.Set("min_trade_count", _minTradeCount)
    return nil
}

// MinTradeCount Getter
func (r TaobaoCrmMembersGetRequest) GetMinTradeCount() int64 {
    return r._minTradeCount
}
// MaxTradeCount Setter
// 最大交易量
func (r *TaobaoCrmMembersGetRequest) SetMaxTradeCount(_maxTradeCount int64) error {
    r._maxTradeCount = _maxTradeCount
    r.Set("max_trade_count", _maxTradeCount)
    return nil
}

// MaxTradeCount Getter
func (r TaobaoCrmMembersGetRequest) GetMaxTradeCount() int64 {
    return r._maxTradeCount
}
// MinLastTradeTime Setter
// 最早上次交易时间
func (r *TaobaoCrmMembersGetRequest) SetMinLastTradeTime(_minLastTradeTime string) error {
    r._minLastTradeTime = _minLastTradeTime
    r.Set("min_last_trade_time", _minLastTradeTime)
    return nil
}

// MinLastTradeTime Getter
func (r TaobaoCrmMembersGetRequest) GetMinLastTradeTime() string {
    return r._minLastTradeTime
}
// MaxLastTradeTime Setter
// 最迟上次交易时间
func (r *TaobaoCrmMembersGetRequest) SetMaxLastTradeTime(_maxLastTradeTime string) error {
    r._maxLastTradeTime = _maxLastTradeTime
    r.Set("max_last_trade_time", _maxLastTradeTime)
    return nil
}

// MaxLastTradeTime Getter
func (r TaobaoCrmMembersGetRequest) GetMaxLastTradeTime() string {
    return r._maxLastTradeTime
}
// PageSize Setter
// 表示每页显示的会员数量,page_size的最大值不能超过100条,最小值不能低于1，
func (r *TaobaoCrmMembersGetRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoCrmMembersGetRequest) GetPageSize() int64 {
    return r._pageSize
}
// CurrentPage Setter
// 显示第几页的会员，如果输入的页码大于总共的页码数，例如总共10页，但是current_page的值为11，则返回空白页，最小页数为1，最大页数为1000
func (r *TaobaoCrmMembersGetRequest) SetCurrentPage(_currentPage int64) error {
    r._currentPage = _currentPage
    r.Set("current_page", _currentPage)
    return nil
}

// CurrentPage Getter
func (r TaobaoCrmMembersGetRequest) GetCurrentPage() int64 {
    return r._currentPage
}
