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
    _pageSize   int64
    // 页码
    _pageNo   int64
    // 应用收费代码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得该应用的收费代码
    _articleCode   string
    // 收费项目代码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得收费项目代码
    _itemCode   string
    // 到期时间起始值（当start_deadline和end_deadline都不填写时，默认返回最近90天的数据）
    _startDeadline   string
    // 到期时间结束值
    _endDeadline   string
    // 订购记录状态，1=有效 2=过期 空=全部
    _status   int64
    // 是否自动续费，true=自动续费 false=非自动续费 空=全部
    _autosub   bool
    // 是否到期提醒，true=到期提醒 false=非到期提醒 空=全部
    _expireNotice   bool
    // 淘宝会员名
    _nick   string
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
func (r *TaobaoVasSubscSearchRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoVasSubscSearchRequest) GetPageSize() int64 {
    return r._pageSize
}
// PageNo Setter
// 页码
func (r *TaobaoVasSubscSearchRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoVasSubscSearchRequest) GetPageNo() int64 {
    return r._pageNo
}
// ArticleCode Setter
// 应用收费代码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得该应用的收费代码
func (r *TaobaoVasSubscSearchRequest) SetArticleCode(_articleCode string) error {
    r._articleCode = _articleCode
    r.Set("article_code", _articleCode)
    return nil
}

// ArticleCode Getter
func (r TaobaoVasSubscSearchRequest) GetArticleCode() string {
    return r._articleCode
}
// ItemCode Setter
// 收费项目代码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得收费项目代码
func (r *TaobaoVasSubscSearchRequest) SetItemCode(_itemCode string) error {
    r._itemCode = _itemCode
    r.Set("item_code", _itemCode)
    return nil
}

// ItemCode Getter
func (r TaobaoVasSubscSearchRequest) GetItemCode() string {
    return r._itemCode
}
// StartDeadline Setter
// 到期时间起始值（当start_deadline和end_deadline都不填写时，默认返回最近90天的数据）
func (r *TaobaoVasSubscSearchRequest) SetStartDeadline(_startDeadline string) error {
    r._startDeadline = _startDeadline
    r.Set("start_deadline", _startDeadline)
    return nil
}

// StartDeadline Getter
func (r TaobaoVasSubscSearchRequest) GetStartDeadline() string {
    return r._startDeadline
}
// EndDeadline Setter
// 到期时间结束值
func (r *TaobaoVasSubscSearchRequest) SetEndDeadline(_endDeadline string) error {
    r._endDeadline = _endDeadline
    r.Set("end_deadline", _endDeadline)
    return nil
}

// EndDeadline Getter
func (r TaobaoVasSubscSearchRequest) GetEndDeadline() string {
    return r._endDeadline
}
// Status Setter
// 订购记录状态，1=有效 2=过期 空=全部
func (r *TaobaoVasSubscSearchRequest) SetStatus(_status int64) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r TaobaoVasSubscSearchRequest) GetStatus() int64 {
    return r._status
}
// Autosub Setter
// 是否自动续费，true=自动续费 false=非自动续费 空=全部
func (r *TaobaoVasSubscSearchRequest) SetAutosub(_autosub bool) error {
    r._autosub = _autosub
    r.Set("autosub", _autosub)
    return nil
}

// Autosub Getter
func (r TaobaoVasSubscSearchRequest) GetAutosub() bool {
    return r._autosub
}
// ExpireNotice Setter
// 是否到期提醒，true=到期提醒 false=非到期提醒 空=全部
func (r *TaobaoVasSubscSearchRequest) SetExpireNotice(_expireNotice bool) error {
    r._expireNotice = _expireNotice
    r.Set("expire_notice", _expireNotice)
    return nil
}

// ExpireNotice Getter
func (r TaobaoVasSubscSearchRequest) GetExpireNotice() bool {
    return r._expireNotice
}
// Nick Setter
// 淘宝会员名
func (r *TaobaoVasSubscSearchRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoVasSubscSearchRequest) GetNick() string {
    return r._nick
}
