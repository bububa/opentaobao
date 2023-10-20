package alihealthmedical

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthmedicalorderrefuseAPIRequest 三方机构通知平台"医生拒诊" API请求
// alibaba.alihealth.medical.order.refuse
//
// 三方机构通知平台&#34;医生拒诊&#34;
type AlibabaalihealthmedicalorderrefuseAPIRequest struct {
	model.Params
	// 请求入参
	_requestInfo *RefuseOrderRequestDto
}

// NewAlibabaalihealthmedicalorderrefuseRequest 初始化AlibabaalihealthmedicalorderrefuseAPIRequest对象
func NewAlibabaalihealthmedicalorderrefuseRequest() *AlibabaalihealthmedicalorderrefuseAPIRequest {
	return &AlibabaalihealthmedicalorderrefuseAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthmedicalorderrefuseAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.medical.order.refuse"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthmedicalorderrefuseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthmedicalorderrefuseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequestInfo is RequestInfo Setter
// 请求入参
func (r *AlibabaalihealthmedicalorderrefuseAPIRequest) SetRequestInfo(_requestInfo *RefuseOrderRequestDto) error {
	r._requestInfo = _requestInfo
	r.Set("request_info", _requestInfo)
	return nil
}

// GetRequestInfo RequestInfo Getter
func (r AlibabaalihealthmedicalorderrefuseAPIRequest) GetRequestInfo() *RefuseOrderRequestDto {
	return r._requestInfo
}
