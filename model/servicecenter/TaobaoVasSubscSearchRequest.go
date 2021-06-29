package servicecenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
订购记录导出 API请求
taobao.vas.subsc.search

用于ISV查询自己名下的应用及收费项目的订购记录
*/
type TaobaoVasSubscSearchRequest struct {
    model.Params
    // 一页包含的记录数
    pageSize   int64
    // 页码
    pageNo   int64
    // 应用收费代码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得该应用的收费代码
    articleCode   string
    // 收费项目代码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得收费项目代码
    itemCode   string
    // 到期时间起始值（当start_deadline和end_deadline都不填写时，默认返回最近90天的数据）
    startDeadline   string
    // 到期时间结束值
    endDeadline   string
    // 订购记录状态，1=有效 2=过期 空=全部
    status   int64
    // 是否自动续费，true=自动续费 false=非自动续费 空=全部
    autosub   bool
    // 是否到期提醒，true=到期提醒 false=非到期提醒 空=全部
    expireNotice   bool
    // 淘宝会员名
    nick   string
}

// 初始化TaobaoVasSubscSearchRequest对象
func NewTaobaoVasSubscSearchRequest() *TaobaoVasSubscSearchRequest{
    return &TaobaoVasSubscSearchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoVasSubscSearchRequest) GetApiMethodName() string {
    return "taobao.vas.subsc.search"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoVasSubscSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PageSize Setter
// 一页包含的记录数
func (r *TaobaoVasSubscSearchRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoVasSubscSearchRequest) GetPageSize() int64 {
    return r.pageSize
}
// PageNo Setter
// 页码
func (r *TaobaoVasSubscSearchRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoVasSubscSearchRequest) GetPageNo() int64 {
    return r.pageNo
}
// ArticleCode Setter
// 应用收费代码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得该应用的收费代码
func (r *TaobaoVasSubscSearchRequest) SetArticleCode(articleCode string) error {
    r.articleCode = articleCode
    r.Set("article_code", articleCode)
    return nil
}

// ArticleCode Getter
func (r TaobaoVasSubscSearchRequest) GetArticleCode() string {
    return r.articleCode
}
// ItemCode Setter
// 收费项目代码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得收费项目代码
func (r *TaobaoVasSubscSearchRequest) SetItemCode(itemCode string) error {
    r.itemCode = itemCode
    r.Set("item_code", itemCode)
    return nil
}

// ItemCode Getter
func (r TaobaoVasSubscSearchRequest) GetItemCode() string {
    return r.itemCode
}
// StartDeadline Setter
// 到期时间起始值（当start_deadline和end_deadline都不填写时，默认返回最近90天的数据）
func (r *TaobaoVasSubscSearchRequest) SetStartDeadline(startDeadline string) error {
    r.startDeadline = startDeadline
    r.Set("start_deadline", startDeadline)
    return nil
}

// StartDeadline Getter
func (r TaobaoVasSubscSearchRequest) GetStartDeadline() string {
    return r.startDeadline
}
// EndDeadline Setter
// 到期时间结束值
func (r *TaobaoVasSubscSearchRequest) SetEndDeadline(endDeadline string) error {
    r.endDeadline = endDeadline
    r.Set("end_deadline", endDeadline)
    return nil
}

// EndDeadline Getter
func (r TaobaoVasSubscSearchRequest) GetEndDeadline() string {
    return r.endDeadline
}
// Status Setter
// 订购记录状态，1=有效 2=过期 空=全部
func (r *TaobaoVasSubscSearchRequest) SetStatus(status int64) error {
    r.status = status
    r.Set("status", status)
    return nil
}

// Status Getter
func (r TaobaoVasSubscSearchRequest) GetStatus() int64 {
    return r.status
}
// Autosub Setter
// 是否自动续费，true=自动续费 false=非自动续费 空=全部
func (r *TaobaoVasSubscSearchRequest) SetAutosub(autosub bool) error {
    r.autosub = autosub
    r.Set("autosub", autosub)
    return nil
}

// Autosub Getter
func (r TaobaoVasSubscSearchRequest) GetAutosub() bool {
    return r.autosub
}
// ExpireNotice Setter
// 是否到期提醒，true=到期提醒 false=非到期提醒 空=全部
func (r *TaobaoVasSubscSearchRequest) SetExpireNotice(expireNotice bool) error {
    r.expireNotice = expireNotice
    r.Set("expire_notice", expireNotice)
    return nil
}

// ExpireNotice Getter
func (r TaobaoVasSubscSearchRequest) GetExpireNotice() bool {
    return r.expireNotice
}
// Nick Setter
// 淘宝会员名
func (r *TaobaoVasSubscSearchRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

// Nick Getter
func (r TaobaoVasSubscSearchRequest) GetNick() string {
    return r.nick
}
