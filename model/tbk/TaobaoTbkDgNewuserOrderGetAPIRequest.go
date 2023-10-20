package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotbkdgnewuserordergetAPIRequest 淘宝客-推广者-新用户订单明细查询 API请求
// taobao.tbk.dg.newuser.order.get
//
// 拉新API
type TaobaotbkdgnewuserordergetAPIRequest struct {
	model.Params
	// 开始时间，当活动为淘宝活动，表示最早注册时间；当活动为支付宝活动，表示最早绑定时间；当活动为天猫活动，表示最早领取红包时间
	_starttime string
	// 结束时间，当活动为淘宝活动，表示最晚结束时间；当活动为支付宝活动，表示最晚绑定时间；当活动为天猫活动，表示最晚领取红包的时间
	_endtime string
	// 活动id， 活动名称与活动ID列表（该字段已废弃）
	_activityid string
	// 页大小，默认20，1~100
	_pagesize int64
	// 页码，默认1
	_pageno int64
	// mm_xxx_xxx_xxx的第三位
	_adzoneid int64
}

// NewTaobaotbkdgnewuserordergetRequest 初始化TaobaotbkdgnewuserordergetAPIRequest对象
func NewTaobaotbkdgnewuserordergetRequest() *TaobaotbkdgnewuserordergetAPIRequest {
	return &TaobaotbkdgnewuserordergetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotbkdgnewuserordergetAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.dg.newuser.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotbkdgnewuserordergetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotbkdgnewuserordergetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStarttime is Starttime Setter
// 开始时间，当活动为淘宝活动，表示最早注册时间；当活动为支付宝活动，表示最早绑定时间；当活动为天猫活动，表示最早领取红包时间
func (r *TaobaotbkdgnewuserordergetAPIRequest) SetStarttime(_starttime string) error {
	r._starttime = _starttime
	r.Set("start_time", _starttime)
	return nil
}

// GetStarttime Starttime Getter
func (r TaobaotbkdgnewuserordergetAPIRequest) GetStarttime() string {
	return r._starttime
}

// SetEndtime is Endtime Setter
// 结束时间，当活动为淘宝活动，表示最晚结束时间；当活动为支付宝活动，表示最晚绑定时间；当活动为天猫活动，表示最晚领取红包的时间
func (r *TaobaotbkdgnewuserordergetAPIRequest) SetEndtime(_endtime string) error {
	r._endtime = _endtime
	r.Set("end_time", _endtime)
	return nil
}

// GetEndtime Endtime Getter
func (r TaobaotbkdgnewuserordergetAPIRequest) GetEndtime() string {
	return r._endtime
}

// SetActivityid is Activityid Setter
// 活动id， 活动名称与活动ID列表（该字段已废弃）
func (r *TaobaotbkdgnewuserordergetAPIRequest) SetActivityid(_activityid string) error {
	r._activityid = _activityid
	r.Set("activity_id", _activityid)
	return nil
}

// GetActivityid Activityid Getter
func (r TaobaotbkdgnewuserordergetAPIRequest) GetActivityid() string {
	return r._activityid
}

// SetPagesize is Pagesize Setter
// 页大小，默认20，1~100
func (r *TaobaotbkdgnewuserordergetAPIRequest) SetPagesize(_pagesize int64) error {
	r._pagesize = _pagesize
	r.Set("page_size", _pagesize)
	return nil
}

// GetPagesize Pagesize Getter
func (r TaobaotbkdgnewuserordergetAPIRequest) GetPagesize() int64 {
	return r._pagesize
}

// SetPageno is Pageno Setter
// 页码，默认1
func (r *TaobaotbkdgnewuserordergetAPIRequest) SetPageno(_pageno int64) error {
	r._pageno = _pageno
	r.Set("page_no", _pageno)
	return nil
}

// GetPageno Pageno Getter
func (r TaobaotbkdgnewuserordergetAPIRequest) GetPageno() int64 {
	return r._pageno
}

// SetAdzoneid is Adzoneid Setter
// mm_xxx_xxx_xxx的第三位
func (r *TaobaotbkdgnewuserordergetAPIRequest) SetAdzoneid(_adzoneid int64) error {
	r._adzoneid = _adzoneid
	r.Set("adzone_id", _adzoneid)
	return nil
}

// GetAdzoneid Adzoneid Getter
func (r TaobaotbkdgnewuserordergetAPIRequest) GetAdzoneid() int64 {
	return r._adzoneid
}
