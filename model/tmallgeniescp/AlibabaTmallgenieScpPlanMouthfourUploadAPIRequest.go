package tmallgeniescp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatmallgeniescpplanmouthfouruploadAPIRequest 21-M+4PR 回传接口接口 API请求
// alibaba.tmallgenie.scp.plan.mouthfour.upload
//
// M+4 PR 回传接口
type AlibabatmallgeniescpplanmouthfouruploadAPIRequest struct {
	model.Params
	// 请求参数
	_monthFourPrRequest *MonthFourPrRequest
}

// NewAlibabatmallgeniescpplanmouthfouruploadRequest 初始化AlibabatmallgeniescpplanmouthfouruploadAPIRequest对象
func NewAlibabatmallgeniescpplanmouthfouruploadRequest() *AlibabatmallgeniescpplanmouthfouruploadAPIRequest {
	return &AlibabatmallgeniescpplanmouthfouruploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatmallgeniescpplanmouthfouruploadAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.mouthfour.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatmallgeniescpplanmouthfouruploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatmallgeniescpplanmouthfouruploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMonthFourPrRequest is MonthFourPrRequest Setter
// 请求参数
func (r *AlibabatmallgeniescpplanmouthfouruploadAPIRequest) SetMonthFourPrRequest(_monthFourPrRequest *MonthFourPrRequest) error {
	r._monthFourPrRequest = _monthFourPrRequest
	r.Set("month_four_pr_request", _monthFourPrRequest)
	return nil
}

// GetMonthFourPrRequest MonthFourPrRequest Getter
func (r AlibabatmallgeniescpplanmouthfouruploadAPIRequest) GetMonthFourPrRequest() *MonthFourPrRequest {
	return r._monthFourPrRequest
}
