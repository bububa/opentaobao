package alihealthpw

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthPwSpecialSynchronodeAPIRequest 合作方同步状态至阿里健康 API请求
// alibaba.alihealth.pw.special.synchronode
//
// 合作方同步状态至阿里健康
type AlibabaAlihealthPwSpecialSynchronodeAPIRequest struct {
	model.Params
	// 状态信息入参
	_body *SNodeDto
}

// NewAlibabaAlihealthPwSpecialSynchronodeRequest 初始化AlibabaAlihealthPwSpecialSynchronodeAPIRequest对象
func NewAlibabaAlihealthPwSpecialSynchronodeRequest() *AlibabaAlihealthPwSpecialSynchronodeAPIRequest {
	return &AlibabaAlihealthPwSpecialSynchronodeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthPwSpecialSynchronodeAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.pw.special.synchronode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthPwSpecialSynchronodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthPwSpecialSynchronodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBody is Body Setter
// 状态信息入参
func (r *AlibabaAlihealthPwSpecialSynchronodeAPIRequest) SetBody(_body *SNodeDto) error {
	r._body = _body
	r.Set("body", _body)
	return nil
}

// GetBody Body Getter
func (r AlibabaAlihealthPwSpecialSynchronodeAPIRequest) GetBody() *SNodeDto {
	return r._body
}
