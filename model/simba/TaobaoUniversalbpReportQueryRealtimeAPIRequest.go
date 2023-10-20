package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpReportQueryRealtimeAPIRequest 实时报表查询 API请求
// taobao.universalbp.report.query.realtime
//
// 实时报表查询
type TaobaoUniversalbpReportQueryRealtimeAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
	// topRealTimeReportQueryVO
	_topRealTimeReportQueryVO *TopRealTimeReportQueryVo
}

// NewTaobaoUniversalbpReportQueryRealtimeRequest 初始化TaobaoUniversalbpReportQueryRealtimeAPIRequest对象
func NewTaobaoUniversalbpReportQueryRealtimeRequest() *TaobaoUniversalbpReportQueryRealtimeAPIRequest {
	return &TaobaoUniversalbpReportQueryRealtimeAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoUniversalbpReportQueryRealtimeAPIRequest) Reset() {
	r._topServiceContext = nil
	r._topRealTimeReportQueryVO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUniversalbpReportQueryRealtimeAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.report.query.realtime"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUniversalbpReportQueryRealtimeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUniversalbpReportQueryRealtimeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaoUniversalbpReportQueryRealtimeAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaoUniversalbpReportQueryRealtimeAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}

// SetTopRealTimeReportQueryVO is TopRealTimeReportQueryVO Setter
// topRealTimeReportQueryVO
func (r *TaobaoUniversalbpReportQueryRealtimeAPIRequest) SetTopRealTimeReportQueryVO(_topRealTimeReportQueryVO *TopRealTimeReportQueryVo) error {
	r._topRealTimeReportQueryVO = _topRealTimeReportQueryVO
	r.Set("top_real_time_report_query_v_o", _topRealTimeReportQueryVO)
	return nil
}

// GetTopRealTimeReportQueryVO TopRealTimeReportQueryVO Getter
func (r TaobaoUniversalbpReportQueryRealtimeAPIRequest) GetTopRealTimeReportQueryVO() *TopRealTimeReportQueryVo {
	return r._topRealTimeReportQueryVO
}

var poolTaobaoUniversalbpReportQueryRealtimeAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoUniversalbpReportQueryRealtimeRequest()
	},
}

// GetTaobaoUniversalbpReportQueryRealtimeRequest 从 sync.Pool 获取 TaobaoUniversalbpReportQueryRealtimeAPIRequest
func GetTaobaoUniversalbpReportQueryRealtimeAPIRequest() *TaobaoUniversalbpReportQueryRealtimeAPIRequest {
	return poolTaobaoUniversalbpReportQueryRealtimeAPIRequest.Get().(*TaobaoUniversalbpReportQueryRealtimeAPIRequest)
}

// ReleaseTaobaoUniversalbpReportQueryRealtimeAPIRequest 将 TaobaoUniversalbpReportQueryRealtimeAPIRequest 放入 sync.Pool
func ReleaseTaobaoUniversalbpReportQueryRealtimeAPIRequest(v *TaobaoUniversalbpReportQueryRealtimeAPIRequest) {
	v.Reset()
	poolTaobaoUniversalbpReportQueryRealtimeAPIRequest.Put(v)
}
