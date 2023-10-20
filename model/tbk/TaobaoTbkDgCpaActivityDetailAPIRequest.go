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
	_indicatorAlias string
	// 指定数据批次号(时间戳)
	_runtime string
	// 明细类型，1：预估明细，2：结算明细
	_queryType int64
	// 每页条数
	_pageSize int64
	// 页码
	_pageNo int64
	// CPA活动ID
	_eventId int64
	// 下一页开始查询的记录主键id
	_startId int64
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

// SetIndicatorAlias is IndicatorAlias Setter
// CPA活动奖励的统计口径，相关说明见文档：https://www.yuque.com/docs/share/7ecf8cf1-7f99-4633-a2ed-f9b6f8116af5?#
func (r *TaobaotbkdgcpaactivitydetailAPIRequest) SetIndicatorAlias(_indicatorAlias string) error {
	r._indicatorAlias = _indicatorAlias
	r.Set("indicator_alias", _indicatorAlias)
	return nil
}

// GetIndicatorAlias IndicatorAlias Getter
func (r TaobaotbkdgcpaactivitydetailAPIRequest) GetIndicatorAlias() string {
	return r._indicatorAlias
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

// SetQueryType is QueryType Setter
// 明细类型，1：预估明细，2：结算明细
func (r *TaobaotbkdgcpaactivitydetailAPIRequest) SetQueryType(_queryType int64) error {
	r._queryType = _queryType
	r.Set("query_type", _queryType)
	return nil
}

// GetQueryType QueryType Getter
func (r TaobaotbkdgcpaactivitydetailAPIRequest) GetQueryType() int64 {
	return r._queryType
}

// SetPageSize is PageSize Setter
// 每页条数
func (r *TaobaotbkdgcpaactivitydetailAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaotbkdgcpaactivitydetailAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNo is PageNo Setter
// 页码
func (r *TaobaotbkdgcpaactivitydetailAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaotbkdgcpaactivitydetailAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetEventId is EventId Setter
// CPA活动ID
func (r *TaobaotbkdgcpaactivitydetailAPIRequest) SetEventId(_eventId int64) error {
	r._eventId = _eventId
	r.Set("event_id", _eventId)
	return nil
}

// GetEventId EventId Getter
func (r TaobaotbkdgcpaactivitydetailAPIRequest) GetEventId() int64 {
	return r._eventId
}

// SetStartId is StartId Setter
// 下一页开始查询的记录主键id
func (r *TaobaotbkdgcpaactivitydetailAPIRequest) SetStartId(_startId int64) error {
	r._startId = _startId
	r.Set("start_id", _startId)
	return nil
}

// GetStartId StartId Getter
func (r TaobaotbkdgcpaactivitydetailAPIRequest) GetStartId() int64 {
	return r._startId
}
