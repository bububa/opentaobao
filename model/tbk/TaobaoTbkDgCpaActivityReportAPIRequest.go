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
	_bizDate string
	// 媒体三段式id（如果传入pid则返回pid汇总数据，不传则返回member维度统计数据，pid和relationid不可同时传入）
	_pid string
	// CPA活动ID，详见https://www.yuque.com/docs/share/16905f3f-3a22-4e7c-b8c3-4d23791d42f7?#
	_eventId int64
	// 分页页数，从1开始
	_pageNo int64
	// 数据类型：数据类型:1预估 2结算 （选择1可查询含当天实时预估统计的累计数据，选择2可查询最晚截止昨天结算的累计数据，具体逻辑以活动规则描述为准；）
	_queryType int64
	// 分页大小
	_pageSize int64
	// 代理id（如果传入rid则返回rid统计数据，不传则返回member维度统计数据）
	_relationId int64
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

// SetBizDate is BizDate Setter
// 日期(yyyyMMdd)
func (r *TaobaotbkdgcpaactivityreportAPIRequest) SetBizDate(_bizDate string) error {
	r._bizDate = _bizDate
	r.Set("biz_date", _bizDate)
	return nil
}

// GetBizDate BizDate Getter
func (r TaobaotbkdgcpaactivityreportAPIRequest) GetBizDate() string {
	return r._bizDate
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

// SetEventId is EventId Setter
// CPA活动ID，详见https://www.yuque.com/docs/share/16905f3f-3a22-4e7c-b8c3-4d23791d42f7?#
func (r *TaobaotbkdgcpaactivityreportAPIRequest) SetEventId(_eventId int64) error {
	r._eventId = _eventId
	r.Set("event_id", _eventId)
	return nil
}

// GetEventId EventId Getter
func (r TaobaotbkdgcpaactivityreportAPIRequest) GetEventId() int64 {
	return r._eventId
}

// SetPageNo is PageNo Setter
// 分页页数，从1开始
func (r *TaobaotbkdgcpaactivityreportAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaotbkdgcpaactivityreportAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetQueryType is QueryType Setter
// 数据类型：数据类型:1预估 2结算 （选择1可查询含当天实时预估统计的累计数据，选择2可查询最晚截止昨天结算的累计数据，具体逻辑以活动规则描述为准；）
func (r *TaobaotbkdgcpaactivityreportAPIRequest) SetQueryType(_queryType int64) error {
	r._queryType = _queryType
	r.Set("query_type", _queryType)
	return nil
}

// GetQueryType QueryType Getter
func (r TaobaotbkdgcpaactivityreportAPIRequest) GetQueryType() int64 {
	return r._queryType
}

// SetPageSize is PageSize Setter
// 分页大小
func (r *TaobaotbkdgcpaactivityreportAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaotbkdgcpaactivityreportAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetRelationId is RelationId Setter
// 代理id（如果传入rid则返回rid统计数据，不传则返回member维度统计数据）
func (r *TaobaotbkdgcpaactivityreportAPIRequest) SetRelationId(_relationId int64) error {
	r._relationId = _relationId
	r.Set("relation_id", _relationId)
	return nil
}

// GetRelationId RelationId Getter
func (r TaobaotbkdgcpaactivityreportAPIRequest) GetRelationId() int64 {
	return r._relationId
}
