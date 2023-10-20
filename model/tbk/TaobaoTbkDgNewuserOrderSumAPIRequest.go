package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotbkdgnewuserordersumAPIRequest 淘宝客-推广者-拉新活动对应数据查询 API请求
// taobao.tbk.dg.newuser.order.sum
//
// 拉新活动汇总API
type TaobaotbkdgnewuserordersumAPIRequest struct {
	model.Params
	// 活动id， 活动名称与活动ID列表，请参见https://tbk.bbs.taobao.com/detail.html?appId=45301&postId=8599277
	_activityid string
	// 结算月份
	_settlemonth string
	// 页大小，默认20，1~100
	_pagesize int64
	// mm_xxx_xxx_xxx的第三位
	_adzoneid int64
	// 页码，默认1
	_pageno int64
	// mm_xxx_xxx_xxx的第二位
	_siteid int64
}

// NewTaobaotbkdgnewuserordersumRequest 初始化TaobaotbkdgnewuserordersumAPIRequest对象
func NewTaobaotbkdgnewuserordersumRequest() *TaobaotbkdgnewuserordersumAPIRequest {
	return &TaobaotbkdgnewuserordersumAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotbkdgnewuserordersumAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.dg.newuser.order.sum"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotbkdgnewuserordersumAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotbkdgnewuserordersumAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetActivityid is Activityid Setter
// 活动id， 活动名称与活动ID列表，请参见https://tbk.bbs.taobao.com/detail.html?appId=45301&amp;postId=8599277
func (r *TaobaotbkdgnewuserordersumAPIRequest) SetActivityid(_activityid string) error {
	r._activityid = _activityid
	r.Set("activity_id", _activityid)
	return nil
}

// GetActivityid Activityid Getter
func (r TaobaotbkdgnewuserordersumAPIRequest) GetActivityid() string {
	return r._activityid
}

// SetSettlemonth is Settlemonth Setter
// 结算月份
func (r *TaobaotbkdgnewuserordersumAPIRequest) SetSettlemonth(_settlemonth string) error {
	r._settlemonth = _settlemonth
	r.Set("settle_month", _settlemonth)
	return nil
}

// GetSettlemonth Settlemonth Getter
func (r TaobaotbkdgnewuserordersumAPIRequest) GetSettlemonth() string {
	return r._settlemonth
}

// SetPagesize is Pagesize Setter
// 页大小，默认20，1~100
func (r *TaobaotbkdgnewuserordersumAPIRequest) SetPagesize(_pagesize int64) error {
	r._pagesize = _pagesize
	r.Set("page_size", _pagesize)
	return nil
}

// GetPagesize Pagesize Getter
func (r TaobaotbkdgnewuserordersumAPIRequest) GetPagesize() int64 {
	return r._pagesize
}

// SetAdzoneid is Adzoneid Setter
// mm_xxx_xxx_xxx的第三位
func (r *TaobaotbkdgnewuserordersumAPIRequest) SetAdzoneid(_adzoneid int64) error {
	r._adzoneid = _adzoneid
	r.Set("adzone_id", _adzoneid)
	return nil
}

// GetAdzoneid Adzoneid Getter
func (r TaobaotbkdgnewuserordersumAPIRequest) GetAdzoneid() int64 {
	return r._adzoneid
}

// SetPageno is Pageno Setter
// 页码，默认1
func (r *TaobaotbkdgnewuserordersumAPIRequest) SetPageno(_pageno int64) error {
	r._pageno = _pageno
	r.Set("page_no", _pageno)
	return nil
}

// GetPageno Pageno Getter
func (r TaobaotbkdgnewuserordersumAPIRequest) GetPageno() int64 {
	return r._pageno
}

// SetSiteid is Siteid Setter
// mm_xxx_xxx_xxx的第二位
func (r *TaobaotbkdgnewuserordersumAPIRequest) SetSiteid(_siteid int64) error {
	r._siteid = _siteid
	r.Set("site_id", _siteid)
	return nil
}

// GetSiteid Siteid Getter
func (r TaobaotbkdgnewuserordersumAPIRequest) GetSiteid() int64 {
	return r._siteid
}
