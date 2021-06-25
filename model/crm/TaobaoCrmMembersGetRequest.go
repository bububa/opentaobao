package crm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取卖家的会员（基本查询） APIRequest
taobao.crm.members.get

查询卖家的会员，进行基本的查询，返回符合条件的会员列表
*/
type TaobaoCrmMembersGetRequest struct {
    model.Params

    // 买家的昵称
    buyerNick   string 

    // 会员等级,如果不传入值则默认为全部等级。
    grade   int64 

    // 最小交易额,单位为元
    minTradeAmount   float64 

    // 最大交易额，单位为元
    maxTradeAmount   float64 

    // 最小交易量
    minTradeCount   int64 

    // 最大交易量
    maxTradeCount   int64 

    // 最早上次交易时间
    minLastTradeTime   string 

    // 最迟上次交易时间
    maxLastTradeTime   string 

    // 表示每页显示的会员数量,page_size的最大值不能超过100条,最小值不能低于1，
    pageSize   int64 

    // 显示第几页的会员，如果输入的页码大于总共的页码数，例如总共10页，但是current_page的值为11，则返回空白页，最小页数为1，最大页数为1000
    currentPage   int64 

}

func NewTaobaoCrmMembersGetRequest() *TaobaoCrmMembersGetRequest{
    return &TaobaoCrmMembersGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoCrmMembersGetRequest) GetApiMethodName() string {
    return "taobao.crm.members.get"
}

func (r TaobaoCrmMembersGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoCrmMembersGetRequest) SetBuyerNick(buyerNick string) error {
    r.buyerNick = buyerNick
    r.Set("buyer_nick", buyerNick)
    return nil
}

func (r TaobaoCrmMembersGetRequest) GetBuyerNick() string {
    return r.buyerNick
}

func (r *TaobaoCrmMembersGetRequest) SetGrade(grade int64) error {
    r.grade = grade
    r.Set("grade", grade)
    return nil
}

func (r TaobaoCrmMembersGetRequest) GetGrade() int64 {
    return r.grade
}

func (r *TaobaoCrmMembersGetRequest) SetMinTradeAmount(minTradeAmount float64) error {
    r.minTradeAmount = minTradeAmount
    r.Set("min_trade_amount", minTradeAmount)
    return nil
}

func (r TaobaoCrmMembersGetRequest) GetMinTradeAmount() float64 {
    return r.minTradeAmount
}

func (r *TaobaoCrmMembersGetRequest) SetMaxTradeAmount(maxTradeAmount float64) error {
    r.maxTradeAmount = maxTradeAmount
    r.Set("max_trade_amount", maxTradeAmount)
    return nil
}

func (r TaobaoCrmMembersGetRequest) GetMaxTradeAmount() float64 {
    return r.maxTradeAmount
}

func (r *TaobaoCrmMembersGetRequest) SetMinTradeCount(minTradeCount int64) error {
    r.minTradeCount = minTradeCount
    r.Set("min_trade_count", minTradeCount)
    return nil
}

func (r TaobaoCrmMembersGetRequest) GetMinTradeCount() int64 {
    return r.minTradeCount
}

func (r *TaobaoCrmMembersGetRequest) SetMaxTradeCount(maxTradeCount int64) error {
    r.maxTradeCount = maxTradeCount
    r.Set("max_trade_count", maxTradeCount)
    return nil
}

func (r TaobaoCrmMembersGetRequest) GetMaxTradeCount() int64 {
    return r.maxTradeCount
}

func (r *TaobaoCrmMembersGetRequest) SetMinLastTradeTime(minLastTradeTime string) error {
    r.minLastTradeTime = minLastTradeTime
    r.Set("min_last_trade_time", minLastTradeTime)
    return nil
}

func (r TaobaoCrmMembersGetRequest) GetMinLastTradeTime() string {
    return r.minLastTradeTime
}

func (r *TaobaoCrmMembersGetRequest) SetMaxLastTradeTime(maxLastTradeTime string) error {
    r.maxLastTradeTime = maxLastTradeTime
    r.Set("max_last_trade_time", maxLastTradeTime)
    return nil
}

func (r TaobaoCrmMembersGetRequest) GetMaxLastTradeTime() string {
    return r.maxLastTradeTime
}

func (r *TaobaoCrmMembersGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoCrmMembersGetRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TaobaoCrmMembersGetRequest) SetCurrentPage(currentPage int64) error {
    r.currentPage = currentPage
    r.Set("current_page", currentPage)
    return nil
}

func (r TaobaoCrmMembersGetRequest) GetCurrentPage() int64 {
    return r.currentPage
}

