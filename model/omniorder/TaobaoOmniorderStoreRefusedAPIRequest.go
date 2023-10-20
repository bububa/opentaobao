package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoomniorderstorerefusedAPIRequest Pos端门店拒单 API请求
// taobao.omniorder.store.refused
//
// ISV Pos端门店拒单，通知星盘
type TaobaoomniorderstorerefusedAPIRequest struct {
	model.Params
	// 子订单列表
	_subOrderList []SubOrder
	// 跟踪Id
	_traceId string
	// 淘宝交易主订单ID
	_tid int64
	// ISV的系统时间
	_reportTimestamp int64
}

// NewTaobaoomniorderstorerefusedRequest 初始化TaobaoomniorderstorerefusedAPIRequest对象
func NewTaobaoomniorderstorerefusedRequest() *TaobaoomniorderstorerefusedAPIRequest {
	return &TaobaoomniorderstorerefusedAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoomniorderstorerefusedAPIRequest) GetApiMethodName() string {
	return "taobao.omniorder.store.refused"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoomniorderstorerefusedAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoomniorderstorerefusedAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSubOrderList is SubOrderList Setter
// 子订单列表
func (r *TaobaoomniorderstorerefusedAPIRequest) SetSubOrderList(_subOrderList []SubOrder) error {
	r._subOrderList = _subOrderList
	r.Set("sub_order_list", _subOrderList)
	return nil
}

// GetSubOrderList SubOrderList Getter
func (r TaobaoomniorderstorerefusedAPIRequest) GetSubOrderList() []SubOrder {
	return r._subOrderList
}

// SetTraceId is TraceId Setter
// 跟踪Id
func (r *TaobaoomniorderstorerefusedAPIRequest) SetTraceId(_traceId string) error {
	r._traceId = _traceId
	r.Set("trace_id", _traceId)
	return nil
}

// GetTraceId TraceId Getter
func (r TaobaoomniorderstorerefusedAPIRequest) GetTraceId() string {
	return r._traceId
}

// SetTid is Tid Setter
// 淘宝交易主订单ID
func (r *TaobaoomniorderstorerefusedAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaoomniorderstorerefusedAPIRequest) GetTid() int64 {
	return r._tid
}

// SetReportTimestamp is ReportTimestamp Setter
// ISV的系统时间
func (r *TaobaoomniorderstorerefusedAPIRequest) SetReportTimestamp(_reportTimestamp int64) error {
	r._reportTimestamp = _reportTimestamp
	r.Set("report_timestamp", _reportTimestamp)
	return nil
}

// GetReportTimestamp ReportTimestamp Getter
func (r TaobaoomniorderstorerefusedAPIRequest) GetReportTimestamp() int64 {
	return r._reportTimestamp
}
