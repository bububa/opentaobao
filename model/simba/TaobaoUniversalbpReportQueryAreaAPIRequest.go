package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpReportQueryAreaAPIRequest 地域报表查询 API请求
// taobao.universalbp.report.query.area
//
// 地域报表查询
type TaobaoUniversalbpReportQueryAreaAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
	// topAreaReportQueryVO
	_topAreaReportQueryVO *TopAreaReportQueryVo
}

// NewTaobaoUniversalbpReportQueryAreaRequest 初始化TaobaoUniversalbpReportQueryAreaAPIRequest对象
func NewTaobaoUniversalbpReportQueryAreaRequest() *TaobaoUniversalbpReportQueryAreaAPIRequest {
	return &TaobaoUniversalbpReportQueryAreaAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoUniversalbpReportQueryAreaAPIRequest) Reset() {
	r._topServiceContext = nil
	r._topAreaReportQueryVO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUniversalbpReportQueryAreaAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.report.query.area"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUniversalbpReportQueryAreaAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUniversalbpReportQueryAreaAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaoUniversalbpReportQueryAreaAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaoUniversalbpReportQueryAreaAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}

// SetTopAreaReportQueryVO is TopAreaReportQueryVO Setter
// topAreaReportQueryVO
func (r *TaobaoUniversalbpReportQueryAreaAPIRequest) SetTopAreaReportQueryVO(_topAreaReportQueryVO *TopAreaReportQueryVo) error {
	r._topAreaReportQueryVO = _topAreaReportQueryVO
	r.Set("top_area_report_query_v_o", _topAreaReportQueryVO)
	return nil
}

// GetTopAreaReportQueryVO TopAreaReportQueryVO Getter
func (r TaobaoUniversalbpReportQueryAreaAPIRequest) GetTopAreaReportQueryVO() *TopAreaReportQueryVo {
	return r._topAreaReportQueryVO
}

var poolTaobaoUniversalbpReportQueryAreaAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoUniversalbpReportQueryAreaRequest()
	},
}

// GetTaobaoUniversalbpReportQueryAreaRequest 从 sync.Pool 获取 TaobaoUniversalbpReportQueryAreaAPIRequest
func GetTaobaoUniversalbpReportQueryAreaAPIRequest() *TaobaoUniversalbpReportQueryAreaAPIRequest {
	return poolTaobaoUniversalbpReportQueryAreaAPIRequest.Get().(*TaobaoUniversalbpReportQueryAreaAPIRequest)
}

// ReleaseTaobaoUniversalbpReportQueryAreaAPIRequest 将 TaobaoUniversalbpReportQueryAreaAPIRequest 放入 sync.Pool
func ReleaseTaobaoUniversalbpReportQueryAreaAPIRequest(v *TaobaoUniversalbpReportQueryAreaAPIRequest) {
	v.Reset()
	poolTaobaoUniversalbpReportQueryAreaAPIRequest.Put(v)
}
