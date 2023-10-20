package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotbkdgcpaactivityreportAPIRequest 淘宝客-推广者-任务奖励效果报表 API请求
// taobao.tbk.dg.cpa.activity.report
//
// 提供给媒体使用的cpa活动报表查询api，当前仅试运行媒体可使用
type TaobaotbkdgcpaactivityreportAPIRequest struct {
	model.Params
	// 日期(yyyyMMdd)
	_bizdate string
	// 媒体三段式id（如果传入pid则返回pid汇总数据，不传则返回member维度统计数据，pid和relationid不可同时传入）
	_pid string
	// CPA活动ID，详见https://www.yuque.com/docs/share/16905f3f-3a22-4e7c-b8c3-4d23791d42f7?#
	_eventid int64
	// 分页页数，从1开始
	_pageno int64
	// 数据类型：数据类型:1预估 2结算 （选择1可查询含当天实时预估统计的累计数据，选择2可查询最晚截止昨天结算的累计数据，具体逻辑以活动规则描述为准；）
	_querytype int64
	// 分页大小
	_pagesize int64
	// 代理id（如果传入rid则返回rid统计数据，不传则返回member维度统计数据）
	_relationid int64
}

// NewTaobaotbkdgcpaactivityreportRequest 初始化TaobaotbkdgcpaactivityreportAPIRequest对象
func NewTaobaotbkdgcpaactivityreportRequest() *TaobaotbkdgcpaactivityreportAPIRequest {
	return &TaobaotbkdgcpaactivityreportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotbkdgcpaactivityreportAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.dg.cpa.activity.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotbkdgcpaactivityreportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotbkdgcpaactivityreportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizdate is Bizdate Setter
// 日期(yyyyMMdd)
func (r *TaobaotbkdgcpaactivityreportAPIRequest) SetBizdate(_bizdate string) error {
	r._bizdate = _bizdate
	r.Set("biz_date", _bizdate)
	return nil
}

// GetBizdate Bizdate Getter
func (r TaobaotbkdgcpaactivityreportAPIRequest) GetBizdate() string {
	return r._bizdate
}

// SetPid is Pid Setter
// 媒体三段式id（如果传入pid则返回pid汇总数据，不传则返回member维度统计数据，pid和relationid不可同时传入）
func (r *TaobaotbkdgcpaactivityreportAPIRequest) SetPid(_pid string) error {
	r._pid = _pid
	r.Set("pid", _pid)
	return nil
}

// GetPid Pid Getter
func (r TaobaotbkdgcpaactivityreportAPIRequest) GetPid() string {
	return r._pid
}

// SetEventid is Eventid Setter
// CPA活动ID，详见https://www.yuque.com/docs/share/16905f3f-3a22-4e7c-b8c3-4d23791d42f7?#
func (r *TaobaotbkdgcpaactivityreportAPIRequest) SetEventid(_eventid int64) error {
	r._eventid = _eventid
	r.Set("event_id", _eventid)
	return nil
}

// GetEventid Eventid Getter
func (r TaobaotbkdgcpaactivityreportAPIRequest) GetEventid() int64 {
	return r._eventid
}

// SetPageno is Pageno Setter
// 分页页数，从1开始
func (r *TaobaotbkdgcpaactivityreportAPIRequest) SetPageno(_pageno int64) error {
	r._pageno = _pageno
	r.Set("page_no", _pageno)
	return nil
}

// GetPageno Pageno Getter
func (r TaobaotbkdgcpaactivityreportAPIRequest) GetPageno() int64 {
	return r._pageno
}

// SetQuerytype is Querytype Setter
// 数据类型：数据类型:1预估 2结算 （选择1可查询含当天实时预估统计的累计数据，选择2可查询最晚截止昨天结算的累计数据，具体逻辑以活动规则描述为准；）
func (r *TaobaotbkdgcpaactivityreportAPIRequest) SetQuerytype(_querytype int64) error {
	r._querytype = _querytype
	r.Set("query_type", _querytype)
	return nil
}

// GetQuerytype Querytype Getter
func (r TaobaotbkdgcpaactivityreportAPIRequest) GetQuerytype() int64 {
	return r._querytype
}

// SetPagesize is Pagesize Setter
// 分页大小
func (r *TaobaotbkdgcpaactivityreportAPIRequest) SetPagesize(_pagesize int64) error {
	r._pagesize = _pagesize
	r.Set("page_size", _pagesize)
	return nil
}

// GetPagesize Pagesize Getter
func (r TaobaotbkdgcpaactivityreportAPIRequest) GetPagesize() int64 {
	return r._pagesize
}

// SetRelationid is Relationid Setter
// 代理id（如果传入rid则返回rid统计数据，不传则返回member维度统计数据）
func (r *TaobaotbkdgcpaactivityreportAPIRequest) SetRelationid(_relationid int64) error {
	r._relationid = _relationid
	r.Set("relation_id", _relationid)
	return nil
}

// GetRelationid Relationid Getter
func (r TaobaotbkdgcpaactivityreportAPIRequest) GetRelationid() int64 {
	return r._relationid
}
