package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoVasSubscSearchAPIRequest 订购记录导出 API请求
// taobao.vas.subsc.search
//
// 用于ISV查询自己名下的应用及收费项目的订购记录
type TaobaoVasSubscSearchAPIRequest struct {
	model.Params
	// 应用收费代码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得该应用的收费代码
	_articleCode string
	// 收费项目代码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得收费项目代码
	_itemCode string
	// 到期时间起始值（当start_deadline和end_deadline都不填写时，默认返回最近90天的数据）
	_startDeadline string
	// 到期时间结束值
	_endDeadline string
	// 淘宝会员名
	_nick string
	// 订购记录状态，1=有效 2=过期 空=全部
	_status int64
	// 一页包含的记录数
	_pageSize int64
	// 页码
	_pageNo int64
	// 是否自动续费，true=自动续费 false=非自动续费 空=全部
	_autosub bool
	// 是否到期提醒，true=到期提醒 false=非到期提醒 空=全部
	_expireNotice bool
}

// NewTaobaoVasSubscSearchRequest 初始化TaobaoVasSubscSearchAPIRequest对象
func NewTaobaoVasSubscSearchRequest() *TaobaoVasSubscSearchAPIRequest {
	return &TaobaoVasSubscSearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoVasSubscSearchAPIRequest) GetApiMethodName() string {
	return "taobao.vas.subsc.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoVasSubscSearchAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetArticleCode is ArticleCode Setter
// 应用收费代码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得该应用的收费代码
func (r *TaobaoVasSubscSearchAPIRequest) SetArticleCode(_articleCode string) error {
	r._articleCode = _articleCode
	r.Set("article_code", _articleCode)
	return nil
}

// GetArticleCode ArticleCode Getter
func (r TaobaoVasSubscSearchAPIRequest) GetArticleCode() string {
	return r._articleCode
}

// SetItemCode is ItemCode Setter
// 收费项目代码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得收费项目代码
func (r *TaobaoVasSubscSearchAPIRequest) SetItemCode(_itemCode string) error {
	r._itemCode = _itemCode
	r.Set("item_code", _itemCode)
	return nil
}

// GetItemCode ItemCode Getter
func (r TaobaoVasSubscSearchAPIRequest) GetItemCode() string {
	return r._itemCode
}

// SetStartDeadline is StartDeadline Setter
// 到期时间起始值（当start_deadline和end_deadline都不填写时，默认返回最近90天的数据）
func (r *TaobaoVasSubscSearchAPIRequest) SetStartDeadline(_startDeadline string) error {
	r._startDeadline = _startDeadline
	r.Set("start_deadline", _startDeadline)
	return nil
}

// GetStartDeadline StartDeadline Getter
func (r TaobaoVasSubscSearchAPIRequest) GetStartDeadline() string {
	return r._startDeadline
}

// SetEndDeadline is EndDeadline Setter
// 到期时间结束值
func (r *TaobaoVasSubscSearchAPIRequest) SetEndDeadline(_endDeadline string) error {
	r._endDeadline = _endDeadline
	r.Set("end_deadline", _endDeadline)
	return nil
}

// GetEndDeadline EndDeadline Getter
func (r TaobaoVasSubscSearchAPIRequest) GetEndDeadline() string {
	return r._endDeadline
}

// SetNick is Nick Setter
// 淘宝会员名
func (r *TaobaoVasSubscSearchAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoVasSubscSearchAPIRequest) GetNick() string {
	return r._nick
}

// SetStatus is Status Setter
// 订购记录状态，1=有效 2=过期 空=全部
func (r *TaobaoVasSubscSearchAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaoVasSubscSearchAPIRequest) GetStatus() int64 {
	return r._status
}

// SetPageSize is PageSize Setter
// 一页包含的记录数
func (r *TaobaoVasSubscSearchAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoVasSubscSearchAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNo is PageNo Setter
// 页码
func (r *TaobaoVasSubscSearchAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoVasSubscSearchAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetAutosub is Autosub Setter
// 是否自动续费，true=自动续费 false=非自动续费 空=全部
func (r *TaobaoVasSubscSearchAPIRequest) SetAutosub(_autosub bool) error {
	r._autosub = _autosub
	r.Set("autosub", _autosub)
	return nil
}

// GetAutosub Autosub Getter
func (r TaobaoVasSubscSearchAPIRequest) GetAutosub() bool {
	return r._autosub
}

// SetExpireNotice is ExpireNotice Setter
// 是否到期提醒，true=到期提醒 false=非到期提醒 空=全部
func (r *TaobaoVasSubscSearchAPIRequest) SetExpireNotice(_expireNotice bool) error {
	r._expireNotice = _expireNotice
	r.Set("expire_notice", _expireNotice)
	return nil
}

// GetExpireNotice ExpireNotice Getter
func (r TaobaoVasSubscSearchAPIRequest) GetExpireNotice() bool {
	return r._expireNotice
}
