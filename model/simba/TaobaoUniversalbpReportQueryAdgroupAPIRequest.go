package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpReportQueryAdgroupAPIRequest 单元报表查询 API请求
// taobao.universalbp.report.query.adgroup
//
// 单元报表查询
type TaobaoUniversalbpReportQueryAdgroupAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
	// topAdgroupReportQueryVO
	_topAdgroupReportQueryVO *TopAdgroupReportQueryVo
}

// NewTaobaoUniversalbpReportQueryAdgroupRequest 初始化TaobaoUniversalbpReportQueryAdgroupAPIRequest对象
func NewTaobaoUniversalbpReportQueryAdgroupRequest() *TaobaoUniversalbpReportQueryAdgroupAPIRequest {
	return &TaobaoUniversalbpReportQueryAdgroupAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoUniversalbpReportQueryAdgroupAPIRequest) Reset() {
	r._topServiceContext = nil
	r._topAdgroupReportQueryVO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUniversalbpReportQueryAdgroupAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.report.query.adgroup"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUniversalbpReportQueryAdgroupAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUniversalbpReportQueryAdgroupAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaoUniversalbpReportQueryAdgroupAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaoUniversalbpReportQueryAdgroupAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}

// SetTopAdgroupReportQueryVO is TopAdgroupReportQueryVO Setter
// topAdgroupReportQueryVO
func (r *TaobaoUniversalbpReportQueryAdgroupAPIRequest) SetTopAdgroupReportQueryVO(_topAdgroupReportQueryVO *TopAdgroupReportQueryVo) error {
	r._topAdgroupReportQueryVO = _topAdgroupReportQueryVO
	r.Set("top_adgroup_report_query_v_o", _topAdgroupReportQueryVO)
	return nil
}

// GetTopAdgroupReportQueryVO TopAdgroupReportQueryVO Getter
func (r TaobaoUniversalbpReportQueryAdgroupAPIRequest) GetTopAdgroupReportQueryVO() *TopAdgroupReportQueryVo {
	return r._topAdgroupReportQueryVO
}

var poolTaobaoUniversalbpReportQueryAdgroupAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoUniversalbpReportQueryAdgroupRequest()
	},
}

// GetTaobaoUniversalbpReportQueryAdgroupRequest 从 sync.Pool 获取 TaobaoUniversalbpReportQueryAdgroupAPIRequest
func GetTaobaoUniversalbpReportQueryAdgroupAPIRequest() *TaobaoUniversalbpReportQueryAdgroupAPIRequest {
	return poolTaobaoUniversalbpReportQueryAdgroupAPIRequest.Get().(*TaobaoUniversalbpReportQueryAdgroupAPIRequest)
}

// ReleaseTaobaoUniversalbpReportQueryAdgroupAPIRequest 将 TaobaoUniversalbpReportQueryAdgroupAPIRequest 放入 sync.Pool
func ReleaseTaobaoUniversalbpReportQueryAdgroupAPIRequest(v *TaobaoUniversalbpReportQueryAdgroupAPIRequest) {
	v.Reset()
	poolTaobaoUniversalbpReportQueryAdgroupAPIRequest.Put(v)
}
