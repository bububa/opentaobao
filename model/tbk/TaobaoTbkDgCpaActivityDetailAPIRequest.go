package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotbkdgcpaactivitydetailAPIRequest 淘宝客-推广者-CPA活动执行明细 API请求
// taobao.tbk.dg.cpa.activity.detail
//
// 淘宝客获取CPA活动具体执行效果的明细数据（含预估和结算维度）
type TaobaotbkdgcpaactivitydetailAPIRequest struct {
	model.Params
	// CPA活动奖励的统计口径，相关说明见文档：https://www.yuque.com/docs/share/7ecf8cf1-7f99-4633-a2ed-f9b6f8116af5?#
	_indicatoralias string
	// 指定数据批次号(时间戳)
	_runtime string
	// 明细类型，1：预估明细，2：结算明细
	_querytype int64
	// 每页条数
	_pagesize int64
	// 页码
	_pageno int64
	// CPA活动ID
	_eventid int64
	// 下一页开始查询的记录主键id
	_startid int64
}

// NewTaobaotbkdgcpaactivitydetailRequest 初始化TaobaotbkdgcpaactivitydetailAPIRequest对象
func NewTaobaotbkdgcpaactivitydetailRequest() *TaobaotbkdgcpaactivitydetailAPIRequest {
	return &TaobaotbkdgcpaactivitydetailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotbkdgcpaactivitydetailAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.dg.cpa.activity.detail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotbkdgcpaactivitydetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotbkdgcpaactivitydetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIndicatoralias is Indicatoralias Setter
// CPA活动奖励的统计口径，相关说明见文档：https://www.yuque.com/docs/share/7ecf8cf1-7f99-4633-a2ed-f9b6f8116af5?#
func (r *TaobaotbkdgcpaactivitydetailAPIRequest) SetIndicatoralias(_indicatoralias string) error {
	r._indicatoralias = _indicatoralias
	r.Set("indicator_alias", _indicatoralias)
	return nil
}

// GetIndicatoralias Indicatoralias Getter
func (r TaobaotbkdgcpaactivitydetailAPIRequest) GetIndicatoralias() string {
	return r._indicatoralias
}

// SetRuntime is Runtime Setter
// 指定数据批次号(时间戳)
func (r *TaobaotbkdgcpaactivitydetailAPIRequest) SetRuntime(_runtime string) error {
	r._runtime = _runtime
	r.Set("runtime", _runtime)
	return nil
}

// GetRuntime Runtime Getter
func (r TaobaotbkdgcpaactivitydetailAPIRequest) GetRuntime() string {
	return r._runtime
}

// SetQuerytype is Querytype Setter
// 明细类型，1：预估明细，2：结算明细
func (r *TaobaotbkdgcpaactivitydetailAPIRequest) SetQuerytype(_querytype int64) error {
	r._querytype = _querytype
	r.Set("query_type", _querytype)
	return nil
}

// GetQuerytype Querytype Getter
func (r TaobaotbkdgcpaactivitydetailAPIRequest) GetQuerytype() int64 {
	return r._querytype
}

// SetPagesize is Pagesize Setter
// 每页条数
func (r *TaobaotbkdgcpaactivitydetailAPIRequest) SetPagesize(_pagesize int64) error {
	r._pagesize = _pagesize
	r.Set("page_size", _pagesize)
	return nil
}

// GetPagesize Pagesize Getter
func (r TaobaotbkdgcpaactivitydetailAPIRequest) GetPagesize() int64 {
	return r._pagesize
}

// SetPageno is Pageno Setter
// 页码
func (r *TaobaotbkdgcpaactivitydetailAPIRequest) SetPageno(_pageno int64) error {
	r._pageno = _pageno
	r.Set("page_no", _pageno)
	return nil
}

// GetPageno Pageno Getter
func (r TaobaotbkdgcpaactivitydetailAPIRequest) GetPageno() int64 {
	return r._pageno
}

// SetEventid is Eventid Setter
// CPA活动ID
func (r *TaobaotbkdgcpaactivitydetailAPIRequest) SetEventid(_eventid int64) error {
	r._eventid = _eventid
	r.Set("event_id", _eventid)
	return nil
}

// GetEventid Eventid Getter
func (r TaobaotbkdgcpaactivitydetailAPIRequest) GetEventid() int64 {
	return r._eventid
}

// SetStartid is Startid Setter
// 下一页开始查询的记录主键id
func (r *TaobaotbkdgcpaactivitydetailAPIRequest) SetStartid(_startid int64) error {
	r._startid = _startid
	r.Set("start_id", _startid)
	return nil
}

// GetStartid Startid Getter
func (r TaobaotbkdgcpaactivitydetailAPIRequest) GetStartid() int64 {
	return r._startid
}
