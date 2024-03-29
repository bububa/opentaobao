package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpReportQueryCrowdAPIRequest 人群报表查询 API请求
// taobao.universalbp.report.query.crowd
//
// 人群报表查询
type TaobaoUniversalbpReportQueryCrowdAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
	// topCrowdReportQueryVO
	_topCrowdReportQueryVO *TopCrowdReportQueryVo
}

// NewTaobaoUniversalbpReportQueryCrowdRequest 初始化TaobaoUniversalbpReportQueryCrowdAPIRequest对象
func NewTaobaoUniversalbpReportQueryCrowdRequest() *TaobaoUniversalbpReportQueryCrowdAPIRequest {
	return &TaobaoUniversalbpReportQueryCrowdAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoUniversalbpReportQueryCrowdAPIRequest) Reset() {
	r._topServiceContext = nil
	r._topCrowdReportQueryVO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUniversalbpReportQueryCrowdAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.report.query.crowd"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUniversalbpReportQueryCrowdAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUniversalbpReportQueryCrowdAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaoUniversalbpReportQueryCrowdAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaoUniversalbpReportQueryCrowdAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}

// SetTopCrowdReportQueryVO is TopCrowdReportQueryVO Setter
// topCrowdReportQueryVO
func (r *TaobaoUniversalbpReportQueryCrowdAPIRequest) SetTopCrowdReportQueryVO(_topCrowdReportQueryVO *TopCrowdReportQueryVo) error {
	r._topCrowdReportQueryVO = _topCrowdReportQueryVO
	r.Set("top_crowd_report_query_v_o", _topCrowdReportQueryVO)
	return nil
}

// GetTopCrowdReportQueryVO TopCrowdReportQueryVO Getter
func (r TaobaoUniversalbpReportQueryCrowdAPIRequest) GetTopCrowdReportQueryVO() *TopCrowdReportQueryVo {
	return r._topCrowdReportQueryVO
}

var poolTaobaoUniversalbpReportQueryCrowdAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoUniversalbpReportQueryCrowdRequest()
	},
}

// GetTaobaoUniversalbpReportQueryCrowdRequest 从 sync.Pool 获取 TaobaoUniversalbpReportQueryCrowdAPIRequest
func GetTaobaoUniversalbpReportQueryCrowdAPIRequest() *TaobaoUniversalbpReportQueryCrowdAPIRequest {
	return poolTaobaoUniversalbpReportQueryCrowdAPIRequest.Get().(*TaobaoUniversalbpReportQueryCrowdAPIRequest)
}

// ReleaseTaobaoUniversalbpReportQueryCrowdAPIRequest 将 TaobaoUniversalbpReportQueryCrowdAPIRequest 放入 sync.Pool
func ReleaseTaobaoUniversalbpReportQueryCrowdAPIRequest(v *TaobaoUniversalbpReportQueryCrowdAPIRequest) {
	v.Reset()
	poolTaobaoUniversalbpReportQueryCrowdAPIRequest.Put(v)
}
