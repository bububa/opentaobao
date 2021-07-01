package crm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取卖家会员（高级查询） API请求
taobao.crm.members.search

会员列表的高级查询，接口返回符合条件的会员列表.<br><br/>注：建议获取09年以后的数据，09年之前的数据不是很完整
*/
type TaobaoCrmMembersSearchAPIRequest struct {
    model.Params
    // 买家昵称
    _buyerNick   string
    // 会员等级
    _grade   int64
    // 最小交易额，单位为元
    _minTradeAmount   float64
    // 最大交易额，单位为元
    _maxTradeAmount   float64
    // 最小交易量
    _minTradeCount   int64
    // 最大交易量
    _maxTradeCount   int64
    // 最早上次交易时间（订单创建时间）
    _minLastTradeTime   string
    // 最迟上次交易时间
    _maxLastTradeTime   string
    // 关系来源，1交易成功，2未成交，3卖家手动吸纳
    _relationSource   int64
    // 分组id
    _groupId   int64
    // 每页显示的会员数量，page_size的最大值不能超过100，最小值不能小于1
    _pageSize   int64
    // 显示第几页的会员，如果输入的页码大于总共的页码数，例如总共10页，但是current_page的值为11，则返回空白页，最小页数为1.最大1000页
    _currentPage   int64
}

// 初始化TaobaoCrmMembersSearchAPIRequest对象
func NewTaobaoCrmMembersSearchRequest() *TaobaoCrmMembersSearchAPIRequest{
    return &TaobaoCrmMembersSearchAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoCrmMembersSearchAPIRequest) GetApiMethodName() string {
    return "taobao.crm.members.search"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoCrmMembersSearchAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BuyerNick Setter
// 买家昵称
func (r *TaobaoCrmMembersSearchAPIRequest) SetBuyerNick(_buyerNick string) error {
    r._buyerNick = _buyerNick
    r.Set("buyer_nick", _buyerNick)
    return nil
}

// BuyerNick Getter
func (r TaobaoCrmMembersSearchAPIRequest) GetBuyerNick() string {
    return r._buyerNick
}
// Grade Setter
// 会员等级
func (r *TaobaoCrmMembersSearchAPIRequest) SetGrade(_grade int64) error {
    r._grade = _grade
    r.Set("grade", _grade)
    return nil
}

// Grade Getter
func (r TaobaoCrmMembersSearchAPIRequest) GetGrade() int64 {
    return r._grade
}
// MinTradeAmount Setter
// 最小交易额，单位为元
func (r *TaobaoCrmMembersSearchAPIRequest) SetMinTradeAmount(_minTradeAmount float64) error {
    r._minTradeAmount = _minTradeAmount
    r.Set("min_trade_amount", _minTradeAmount)
    return nil
}

// MinTradeAmount Getter
func (r TaobaoCrmMembersSearchAPIRequest) GetMinTradeAmount() float64 {
    return r._minTradeAmount
}
// MaxTradeAmount Setter
// 最大交易额，单位为元
func (r *TaobaoCrmMembersSearchAPIRequest) SetMaxTradeAmount(_maxTradeAmount float64) error {
    r._maxTradeAmount = _maxTradeAmount
    r.Set("max_trade_amount", _maxTradeAmount)
    return nil
}

// MaxTradeAmount Getter
func (r TaobaoCrmMembersSearchAPIRequest) GetMaxTradeAmount() float64 {
    return r._maxTradeAmount
}
// MinTradeCount Setter
// 最小交易量
func (r *TaobaoCrmMembersSearchAPIRequest) SetMinTradeCount(_minTradeCount int64) error {
    r._minTradeCount = _minTradeCount
    r.Set("min_trade_count", _minTradeCount)
    return nil
}

// MinTradeCount Getter
func (r TaobaoCrmMembersSearchAPIRequest) GetMinTradeCount() int64 {
    return r._minTradeCount
}
// MaxTradeCount Setter
// 最大交易量
func (r *TaobaoCrmMembersSearchAPIRequest) SetMaxTradeCount(_maxTradeCount int64) error {
    r._maxTradeCount = _maxTradeCount
    r.Set("max_trade_count", _maxTradeCount)
    return nil
}

// MaxTradeCount Getter
func (r TaobaoCrmMembersSearchAPIRequest) GetMaxTradeCount() int64 {
    return r._maxTradeCount
}
// MinLastTradeTime Setter
// 最早上次交易时间（订单创建时间）
func (r *TaobaoCrmMembersSearchAPIRequest) SetMinLastTradeTime(_minLastTradeTime string) error {
    r._minLastTradeTime = _minLastTradeTime
    r.Set("min_last_trade_time", _minLastTradeTime)
    return nil
}

// MinLastTradeTime Getter
func (r TaobaoCrmMembersSearchAPIRequest) GetMinLastTradeTime() string {
    return r._minLastTradeTime
}
// MaxLastTradeTime Setter
// 最迟上次交易时间
func (r *TaobaoCrmMembersSearchAPIRequest) SetMaxLastTradeTime(_maxLastTradeTime string) error {
    r._maxLastTradeTime = _maxLastTradeTime
    r.Set("max_last_trade_time", _maxLastTradeTime)
    return nil
}

// MaxLastTradeTime Getter
func (r TaobaoCrmMembersSearchAPIRequest) GetMaxLastTradeTime() string {
    return r._maxLastTradeTime
}
// RelationSource Setter
// 关系来源，1交易成功，2未成交，3卖家手动吸纳
func (r *TaobaoCrmMembersSearchAPIRequest) SetRelationSource(_relationSource int64) error {
    r._relationSource = _relationSource
    r.Set("relation_source", _relationSource)
    return nil
}

// RelationSource Getter
func (r TaobaoCrmMembersSearchAPIRequest) GetRelationSource() int64 {
    return r._relationSource
}
// GroupId Setter
// 分组id
func (r *TaobaoCrmMembersSearchAPIRequest) SetGroupId(_groupId int64) error {
    r._groupId = _groupId
    r.Set("group_id", _groupId)
    return nil
}

// GroupId Getter
func (r TaobaoCrmMembersSearchAPIRequest) GetGroupId() int64 {
    return r._groupId
}
// PageSize Setter
// 每页显示的会员数量，page_size的最大值不能超过100，最小值不能小于1
func (r *TaobaoCrmMembersSearchAPIRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoCrmMembersSearchAPIRequest) GetPageSize() int64 {
    return r._pageSize
}
// CurrentPage Setter
// 显示第几页的会员，如果输入的页码大于总共的页码数，例如总共10页，但是current_page的值为11，则返回空白页，最小页数为1.最大1000页
func (r *TaobaoCrmMembersSearchAPIRequest) SetCurrentPage(_currentPage int64) error {
    r._currentPage = _currentPage
    r.Set("current_page", _currentPage)
    return nil
}

// CurrentPage Getter
func (r TaobaoCrmMembersSearchAPIRequest) GetCurrentPage() int64 {
    return r._currentPage
}
