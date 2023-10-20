package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaaliqinflowwalletcheckbalanceAPIRequest 商家预存余额检查 API请求
// alibaba.aliqin.flow.wallet.check.balance
//
// 检查商家CRM预存余额是否足够进行活动
type AlibabaaliqinflowwalletcheckbalanceAPIRequest struct {
	model.Params
	// 检查金额档位id
	_gradeId string
}

// NewAlibabaaliqinflowwalletcheckbalanceRequest 初始化AlibabaaliqinflowwalletcheckbalanceAPIRequest对象
func NewAlibabaaliqinflowwalletcheckbalanceRequest() *AlibabaaliqinflowwalletcheckbalanceAPIRequest {
	return &AlibabaaliqinflowwalletcheckbalanceAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaaliqinflowwalletcheckbalanceAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.flow.wallet.check.balance"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaaliqinflowwalletcheckbalanceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaaliqinflowwalletcheckbalanceAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGradeId is GradeId Setter
// 检查金额档位id
func (r *AlibabaaliqinflowwalletcheckbalanceAPIRequest) SetGradeId(_gradeId string) error {
	r._gradeId = _gradeId
	r.Set("grade_id", _gradeId)
	return nil
}

// GetGradeId GradeId Getter
func (r AlibabaaliqinflowwalletcheckbalanceAPIRequest) GetGradeId() string {
	return r._gradeId
}
