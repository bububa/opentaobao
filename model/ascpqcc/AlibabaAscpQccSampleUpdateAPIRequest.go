package ascpqcc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpqccsampleupdateAPIRequest 品控中心更新样品信息 API请求
// alibaba.ascp.qcc.sample.update
//
// 品控中心更新样品信息
type AlibabaascpqccsampleupdateAPIRequest struct {
	model.Params
	// 更新请求参数
	_updateRequest *UpdateSampleRequest
}

// NewAlibabaascpqccsampleupdateRequest 初始化AlibabaascpqccsampleupdateAPIRequest对象
func NewAlibabaascpqccsampleupdateRequest() *AlibabaascpqccsampleupdateAPIRequest {
	return &AlibabaascpqccsampleupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaascpqccsampleupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.qcc.sample.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaascpqccsampleupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaascpqccsampleupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUpdateRequest is UpdateRequest Setter
// 更新请求参数
func (r *AlibabaascpqccsampleupdateAPIRequest) SetUpdateRequest(_updateRequest *UpdateSampleRequest) error {
	r._updateRequest = _updateRequest
	r.Set("update_request", _updateRequest)
	return nil
}

// GetUpdateRequest UpdateRequest Getter
func (r AlibabaascpqccsampleupdateAPIRequest) GetUpdateRequest() *UpdateSampleRequest {
	return r._updateRequest
}
