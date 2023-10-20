package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoomniorderstoreaccpetedAPIRequest Pos端门店接单接口 API请求
// taobao.omniorder.store.accpeted
//
// ISV Pos端门店接单，通知星盘
type TaobaoomniorderstoreaccpetedAPIRequest struct {
	model.Params
	// 子订单列表
	_subOrderList []StoreAcceptedResult
	// 跟踪Id
	_traceId string
	// 淘宝交易主订单ID
	_tid int64
	// ISV系统上报时间
	_reportTimestamp int64
}

// NewTaobaoomniorderstoreaccpetedRequest 初始化TaobaoomniorderstoreaccpetedAPIRequest对象
func NewTaobaoomniorderstoreaccpetedRequest() *TaobaoomniorderstoreaccpetedAPIRequest {
	return &TaobaoomniorderstoreaccpetedAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoomniorderstoreaccpetedAPIRequest) GetApiMethodName() string {
	return "taobao.omniorder.store.accpeted"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoomniorderstoreaccpetedAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoomniorderstoreaccpetedAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSubOrderList is SubOrderList Setter
// 子订单列表
func (r *TaobaoomniorderstoreaccpetedAPIRequest) SetSubOrderList(_subOrderList []StoreAcceptedResult) error {
	r._subOrderList = _subOrderList
	r.Set("sub_order_list", _subOrderList)
	return nil
}

// GetSubOrderList SubOrderList Getter
func (r TaobaoomniorderstoreaccpetedAPIRequest) GetSubOrderList() []StoreAcceptedResult {
	return r._subOrderList
}

// SetTraceId is TraceId Setter
// 跟踪Id
func (r *TaobaoomniorderstoreaccpetedAPIRequest) SetTraceId(_traceId string) error {
	r._traceId = _traceId
	r.Set("trace_id", _traceId)
	return nil
}

// GetTraceId TraceId Getter
func (r TaobaoomniorderstoreaccpetedAPIRequest) GetTraceId() string {
	return r._traceId
}

// SetTid is Tid Setter
// 淘宝交易主订单ID
func (r *TaobaoomniorderstoreaccpetedAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaoomniorderstoreaccpetedAPIRequest) GetTid() int64 {
	return r._tid
}

// SetReportTimestamp is ReportTimestamp Setter
// ISV系统上报时间
func (r *TaobaoomniorderstoreaccpetedAPIRequest) SetReportTimestamp(_reportTimestamp int64) error {
	r._reportTimestamp = _reportTimestamp
	r.Set("report_timestamp", _reportTimestamp)
	return nil
}

// GetReportTimestamp ReportTimestamp Getter
func (r TaobaoomniorderstoreaccpetedAPIRequest) GetReportTimestamp() int64 {
	return r._reportTimestamp
}
