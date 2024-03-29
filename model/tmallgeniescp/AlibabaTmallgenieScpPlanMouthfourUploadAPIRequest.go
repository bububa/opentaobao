package tmallgeniescp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTmallgenieScpPlanMouthfourUploadAPIRequest 21-M+4PR 回传接口接口 API请求
// alibaba.tmallgenie.scp.plan.mouthfour.upload
//
// M+4 PR 回传接口
type AlibabaTmallgenieScpPlanMouthfourUploadAPIRequest struct {
	model.Params
	// 请求参数
	_monthFourPrRequest *MonthFourPrRequest
}

// NewAlibabaTmallgenieScpPlanMouthfourUploadRequest 初始化AlibabaTmallgenieScpPlanMouthfourUploadAPIRequest对象
func NewAlibabaTmallgenieScpPlanMouthfourUploadRequest() *AlibabaTmallgenieScpPlanMouthfourUploadAPIRequest {
	return &AlibabaTmallgenieScpPlanMouthfourUploadAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaTmallgenieScpPlanMouthfourUploadAPIRequest) Reset() {
	r._monthFourPrRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanMouthfourUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.mouthfour.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanMouthfourUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTmallgenieScpPlanMouthfourUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMonthFourPrRequest is MonthFourPrRequest Setter
// 请求参数
func (r *AlibabaTmallgenieScpPlanMouthfourUploadAPIRequest) SetMonthFourPrRequest(_monthFourPrRequest *MonthFourPrRequest) error {
	r._monthFourPrRequest = _monthFourPrRequest
	r.Set("month_four_pr_request", _monthFourPrRequest)
	return nil
}

// GetMonthFourPrRequest MonthFourPrRequest Getter
func (r AlibabaTmallgenieScpPlanMouthfourUploadAPIRequest) GetMonthFourPrRequest() *MonthFourPrRequest {
	return r._monthFourPrRequest
}

var poolAlibabaTmallgenieScpPlanMouthfourUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaTmallgenieScpPlanMouthfourUploadRequest()
	},
}

// GetAlibabaTmallgenieScpPlanMouthfourUploadRequest 从 sync.Pool 获取 AlibabaTmallgenieScpPlanMouthfourUploadAPIRequest
func GetAlibabaTmallgenieScpPlanMouthfourUploadAPIRequest() *AlibabaTmallgenieScpPlanMouthfourUploadAPIRequest {
	return poolAlibabaTmallgenieScpPlanMouthfourUploadAPIRequest.Get().(*AlibabaTmallgenieScpPlanMouthfourUploadAPIRequest)
}

// ReleaseAlibabaTmallgenieScpPlanMouthfourUploadAPIRequest 将 AlibabaTmallgenieScpPlanMouthfourUploadAPIRequest 放入 sync.Pool
func ReleaseAlibabaTmallgenieScpPlanMouthfourUploadAPIRequest(v *AlibabaTmallgenieScpPlanMouthfourUploadAPIRequest) {
	v.Reset()
	poolAlibabaTmallgenieScpPlanMouthfourUploadAPIRequest.Put(v)
}
