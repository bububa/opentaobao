package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotbkscvegassendreportAPIRequest 淘宝客-服务商-查询红包发放个数 API请求
// taobao.tbk.sc.vegas.send.report
//
// 服务商使用。可通过此接口查询对应推广者账号下的红包发放个数。
type TaobaotbkscvegassendreportAPIRequest struct {
	model.Params
	// 统计日期
	_bizdate string
	// 媒体推广pid
	_pid string
	// 查询维度，不填写默认是pid维度
	_rptdim string
	// 渠道关系id
	_relationid int64
	// 已下线，后续不需要填写
	_activityid int64
	// 页码
	_pageno int64
	// 每页大小
	_pagesize int64
	// 查询红包类型，1-超级红包，2-福利购，3-签到红包，4-福利直降，不传时默认查询超级红包数据
	_activitycategory int64
}

// NewTaobaotbkscvegassendreportRequest 初始化TaobaotbkscvegassendreportAPIRequest对象
func NewTaobaotbkscvegassendreportRequest() *TaobaotbkscvegassendreportAPIRequest {
	return &TaobaotbkscvegassendreportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotbkscvegassendreportAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.sc.vegas.send.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotbkscvegassendreportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotbkscvegassendreportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizdate is Bizdate Setter
// 统计日期
func (r *TaobaotbkscvegassendreportAPIRequest) SetBizdate(_bizdate string) error {
	r._bizdate = _bizdate
	r.Set("biz_date", _bizdate)
	return nil
}

// GetBizdate Bizdate Getter
func (r TaobaotbkscvegassendreportAPIRequest) GetBizdate() string {
	return r._bizdate
}

// SetPid is Pid Setter
// 媒体推广pid
func (r *TaobaotbkscvegassendreportAPIRequest) SetPid(_pid string) error {
	r._pid = _pid
	r.Set("pid", _pid)
	return nil
}

// GetPid Pid Getter
func (r TaobaotbkscvegassendreportAPIRequest) GetPid() string {
	return r._pid
}

// SetRptdim is Rptdim Setter
// 查询维度，不填写默认是pid维度
func (r *TaobaotbkscvegassendreportAPIRequest) SetRptdim(_rptdim string) error {
	r._rptdim = _rptdim
	r.Set("rpt_dim", _rptdim)
	return nil
}

// GetRptdim Rptdim Getter
func (r TaobaotbkscvegassendreportAPIRequest) GetRptdim() string {
	return r._rptdim
}

// SetRelationid is Relationid Setter
// 渠道关系id
func (r *TaobaotbkscvegassendreportAPIRequest) SetRelationid(_relationid int64) error {
	r._relationid = _relationid
	r.Set("relation_id", _relationid)
	return nil
}

// GetRelationid Relationid Getter
func (r TaobaotbkscvegassendreportAPIRequest) GetRelationid() int64 {
	return r._relationid
}

// SetActivityid is Activityid Setter
// 已下线，后续不需要填写
func (r *TaobaotbkscvegassendreportAPIRequest) SetActivityid(_activityid int64) error {
	r._activityid = _activityid
	r.Set("activity_id", _activityid)
	return nil
}

// GetActivityid Activityid Getter
func (r TaobaotbkscvegassendreportAPIRequest) GetActivityid() int64 {
	return r._activityid
}

// SetPageno is Pageno Setter
// 页码
func (r *TaobaotbkscvegassendreportAPIRequest) SetPageno(_pageno int64) error {
	r._pageno = _pageno
	r.Set("page_no", _pageno)
	return nil
}

// GetPageno Pageno Getter
func (r TaobaotbkscvegassendreportAPIRequest) GetPageno() int64 {
	return r._pageno
}

// SetPagesize is Pagesize Setter
// 每页大小
func (r *TaobaotbkscvegassendreportAPIRequest) SetPagesize(_pagesize int64) error {
	r._pagesize = _pagesize
	r.Set("page_size", _pagesize)
	return nil
}

// GetPagesize Pagesize Getter
func (r TaobaotbkscvegassendreportAPIRequest) GetPagesize() int64 {
	return r._pagesize
}

// SetActivitycategory is Activitycategory Setter
// 查询红包类型，1-超级红包，2-福利购，3-签到红包，4-福利直降，不传时默认查询超级红包数据
func (r *TaobaotbkscvegassendreportAPIRequest) SetActivitycategory(_activitycategory int64) error {
	r._activitycategory = _activitycategory
	r.Set("activity_category", _activitycategory)
	return nil
}

// GetActivitycategory Activitycategory Getter
func (r TaobaotbkscvegassendreportAPIRequest) GetActivitycategory() int64 {
	return r._activitycategory
}
