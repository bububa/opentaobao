package crm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取卖家会员（高级查询） APIRequest
taobao.crm.members.search

会员列表的高级查询，接口返回符合条件的会员列表.<br><br/>注：建议获取09年以后的数据，09年之前的数据不是很完整
*/
type TaobaoCrmMembersSearchRequest struct {
    model.Params

    // 买家昵称
    buyerNick   string 

    // 会员等级
    grade   int64 

    // 最小交易额，单位为元
    minTradeAmount   float64 

    // 最大交易额，单位为元
    maxTradeAmount   float64 

    // 最小交易量
    minTradeCount   int64 

    // 最大交易量
    maxTradeCount   int64 

    // 最早上次交易时间（订单创建时间）
    minLastTradeTime   string 

    // 最迟上次交易时间
    maxLastTradeTime   string 

    // 关系来源，1交易成功，2未成交，3卖家手动吸纳
    relationSource   int64 

    // 分组id
    groupId   int64 

    // 每页显示的会员数量，page_size的最大值不能超过100，最小值不能小于1
    pageSize   int64 

    // 显示第几页的会员，如果输入的页码大于总共的页码数，例如总共10页，但是current_page的值为11，则返回空白页，最小页数为1.最大1000页
    currentPage   int64 

}

func NewTaobaoCrmMembersSearchRequest() *TaobaoCrmMembersSearchRequest{
    return &TaobaoCrmMembersSearchRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoCrmMembersSearchRequest) GetApiMethodName() string {
    return "taobao.crm.members.search"
}

func (r TaobaoCrmMembersSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoCrmMembersSearchRequest) SetBuyerNick(buyerNick string) error {
    r.buyerNick = buyerNick
    r.Set("buyer_nick", buyerNick)
    return nil
}

func (r TaobaoCrmMembersSearchRequest) GetBuyerNick() string {
    return r.buyerNick
}

func (r *TaobaoCrmMembersSearchRequest) SetGrade(grade int64) error {
    r.grade = grade
    r.Set("grade", grade)
    return nil
}

func (r TaobaoCrmMembersSearchRequest) GetGrade() int64 {
    return r.grade
}

func (r *TaobaoCrmMembersSearchRequest) SetMinTradeAmount(minTradeAmount float64) error {
    r.minTradeAmount = minTradeAmount
    r.Set("min_trade_amount", minTradeAmount)
    return nil
}

func (r TaobaoCrmMembersSearchRequest) GetMinTradeAmount() float64 {
    return r.minTradeAmount
}

func (r *TaobaoCrmMembersSearchRequest) SetMaxTradeAmount(maxTradeAmount float64) error {
    r.maxTradeAmount = maxTradeAmount
    r.Set("max_trade_amount", maxTradeAmount)
    return nil
}

func (r TaobaoCrmMembersSearchRequest) GetMaxTradeAmount() float64 {
    return r.maxTradeAmount
}

func (r *TaobaoCrmMembersSearchRequest) SetMinTradeCount(minTradeCount int64) error {
    r.minTradeCount = minTradeCount
    r.Set("min_trade_count", minTradeCount)
    return nil
}

func (r TaobaoCrmMembersSearchRequest) GetMinTradeCount() int64 {
    return r.minTradeCount
}

func (r *TaobaoCrmMembersSearchRequest) SetMaxTradeCount(maxTradeCount int64) error {
    r.maxTradeCount = maxTradeCount
    r.Set("max_trade_count", maxTradeCount)
    return nil
}

func (r TaobaoCrmMembersSearchRequest) GetMaxTradeCount() int64 {
    return r.maxTradeCount
}

func (r *TaobaoCrmMembersSearchRequest) SetMinLastTradeTime(minLastTradeTime string) error {
    r.minLastTradeTime = minLastTradeTime
    r.Set("min_last_trade_time", minLastTradeTime)
    return nil
}

func (r TaobaoCrmMembersSearchRequest) GetMinLastTradeTime() string {
    return r.minLastTradeTime
}

func (r *TaobaoCrmMembersSearchRequest) SetMaxLastTradeTime(maxLastTradeTime string) error {
    r.maxLastTradeTime = maxLastTradeTime
    r.Set("max_last_trade_time", maxLastTradeTime)
    return nil
}

func (r TaobaoCrmMembersSearchRequest) GetMaxLastTradeTime() string {
    return r.maxLastTradeTime
}

func (r *TaobaoCrmMembersSearchRequest) SetRelationSource(relationSource int64) error {
    r.relationSource = relationSource
    r.Set("relation_source", relationSource)
    return nil
}

func (r TaobaoCrmMembersSearchRequest) GetRelationSource() int64 {
    return r.relationSource
}

func (r *TaobaoCrmMembersSearchRequest) SetGroupId(groupId int64) error {
    r.groupId = groupId
    r.Set("group_id", groupId)
    return nil
}

func (r TaobaoCrmMembersSearchRequest) GetGroupId() int64 {
    return r.groupId
}

func (r *TaobaoCrmMembersSearchRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoCrmMembersSearchRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TaobaoCrmMembersSearchRequest) SetCurrentPage(currentPage int64) error {
    r.currentPage = currentPage
    r.Set("current_page", currentPage)
    return nil
}

func (r TaobaoCrmMembersSearchRequest) GetCurrentPage() int64 {
    return r.currentPage
}

