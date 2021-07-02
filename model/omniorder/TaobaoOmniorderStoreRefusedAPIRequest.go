package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniorderStoreRefusedAPIRequest Pos端门店拒单 API请求
// taobao.omniorder.store.refused
//
// ISV Pos端门店拒单，通知星盘
type TaobaoOmniorderStoreRefusedAPIRequest struct {
	model.Params
	// 淘宝交易主订单ID
	_tid int64
	// 子订单列表
	_subOrderList []SubOrder
	// ISV的系统时间
	_reportTimestamp int64
	// 跟踪Id
	_traceId string
}

// NewTaobaoOmniorderStoreRefusedRequest 初始化TaobaoOmniorderStoreRefusedAPIRequest对象
func NewTaobaoOmniorderStoreRefusedRequest() *TaobaoOmniorderStoreRefusedAPIRequest {
	return &TaobaoOmniorderStoreRefusedAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOmniorderStoreRefusedAPIRequest) GetApiMethodName() string {
	return "taobao.omniorder.store.refused"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOmniorderStoreRefusedAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Tid Setter
// 淘宝交易主订单ID
func (r *TaobaoOmniorderStoreRefusedAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// Get Tid Getter
func (r TaobaoOmniorderStoreRefusedAPIRequest) GetTid() int64 {
	return r._tid
}

// Set is SubOrderList Setter
// 子订单列表
func (r *TaobaoOmniorderStoreRefusedAPIRequest) SetSubOrderList(_subOrderList []SubOrder) error {
	r._subOrderList = _subOrderList
	r.Set("sub_order_list", _subOrderList)
	return nil
}

// Get SubOrderList Getter
func (r TaobaoOmniorderStoreRefusedAPIRequest) GetSubOrderList() []SubOrder {
	return r._subOrderList
}

// Set is ReportTimestamp Setter
// ISV的系统时间
func (r *TaobaoOmniorderStoreRefusedAPIRequest) SetReportTimestamp(_reportTimestamp int64) error {
	r._reportTimestamp = _reportTimestamp
	r.Set("report_timestamp", _reportTimestamp)
	return nil
}

// Get ReportTimestamp Getter
func (r TaobaoOmniorderStoreRefusedAPIRequest) GetReportTimestamp() int64 {
	return r._reportTimestamp
}

// Set is TraceId Setter
// 跟踪Id
func (r *TaobaoOmniorderStoreRefusedAPIRequest) SetTraceId(_traceId string) error {
	r._traceId = _traceId
	r.Set("trace_id", _traceId)
	return nil
}

// Get TraceId Getter
func (r TaobaoOmniorderStoreRefusedAPIRequest) GetTraceId() string {
	return r._traceId
}
