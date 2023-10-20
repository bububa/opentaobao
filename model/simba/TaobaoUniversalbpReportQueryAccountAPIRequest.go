package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpReportQueryAccountAPIRequest 账户报表查询 API请求
// taobao.universalbp.report.query.account
//
// 账户报表查询
type TaobaoUniversalbpReportQueryAccountAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
	// topAccountReportQueryVO
	_topAccountReportQueryVO *TopAccountReportQueryVo
}

// NewTaobaoUniversalbpReportQueryAccountRequest 初始化TaobaoUniversalbpReportQueryAccountAPIRequest对象
func NewTaobaoUniversalbpReportQueryAccountRequest() *TaobaoUniversalbpReportQueryAccountAPIRequest {
	return &TaobaoUniversalbpReportQueryAccountAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoUniversalbpReportQueryAccountAPIRequest) Reset() {
	r._topServiceContext = nil
	r._topAccountReportQueryVO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUniversalbpReportQueryAccountAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.report.query.account"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUniversalbpReportQueryAccountAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUniversalbpReportQueryAccountAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaoUniversalbpReportQueryAccountAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaoUniversalbpReportQueryAccountAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}

// SetTopAccountReportQueryVO is TopAccountReportQueryVO Setter
// topAccountReportQueryVO
func (r *TaobaoUniversalbpReportQueryAccountAPIRequest) SetTopAccountReportQueryVO(_topAccountReportQueryVO *TopAccountReportQueryVo) error {
	r._topAccountReportQueryVO = _topAccountReportQueryVO
	r.Set("top_account_report_query_v_o", _topAccountReportQueryVO)
	return nil
}

// GetTopAccountReportQueryVO TopAccountReportQueryVO Getter
func (r TaobaoUniversalbpReportQueryAccountAPIRequest) GetTopAccountReportQueryVO() *TopAccountReportQueryVo {
	return r._topAccountReportQueryVO
}

var poolTaobaoUniversalbpReportQueryAccountAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoUniversalbpReportQueryAccountRequest()
	},
}

// GetTaobaoUniversalbpReportQueryAccountRequest 从 sync.Pool 获取 TaobaoUniversalbpReportQueryAccountAPIRequest
func GetTaobaoUniversalbpReportQueryAccountAPIRequest() *TaobaoUniversalbpReportQueryAccountAPIRequest {
	return poolTaobaoUniversalbpReportQueryAccountAPIRequest.Get().(*TaobaoUniversalbpReportQueryAccountAPIRequest)
}

// ReleaseTaobaoUniversalbpReportQueryAccountAPIRequest 将 TaobaoUniversalbpReportQueryAccountAPIRequest 放入 sync.Pool
func ReleaseTaobaoUniversalbpReportQueryAccountAPIRequest(v *TaobaoUniversalbpReportQueryAccountAPIRequest) {
	v.Reset()
	poolTaobaoUniversalbpReportQueryAccountAPIRequest.Put(v)
}
