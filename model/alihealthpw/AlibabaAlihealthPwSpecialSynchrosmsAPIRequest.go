package alihealthpw

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthPwSpecialSynchrosmsAPIRequest 同步短信信息至阿里健康 API请求
// alibaba.alihealth.pw.special.synchrosms
//
// 同步短信信息至阿里健康
type AlibabaAlihealthPwSpecialSynchrosmsAPIRequest struct {
	model.Params
	// 入参
	_body *SynchroSmsDto
}

// NewAlibabaAlihealthPwSpecialSynchrosmsRequest 初始化AlibabaAlihealthPwSpecialSynchrosmsAPIRequest对象
func NewAlibabaAlihealthPwSpecialSynchrosmsRequest() *AlibabaAlihealthPwSpecialSynchrosmsAPIRequest {
	return &AlibabaAlihealthPwSpecialSynchrosmsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthPwSpecialSynchrosmsAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.pw.special.synchrosms"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthPwSpecialSynchrosmsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthPwSpecialSynchrosmsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBody is Body Setter
// 入参
func (r *AlibabaAlihealthPwSpecialSynchrosmsAPIRequest) SetBody(_body *SynchroSmsDto) error {
	r._body = _body
	r.Set("body", _body)
	return nil
}

// GetBody Body Getter
func (r AlibabaAlihealthPwSpecialSynchrosmsAPIRequest) GetBody() *SynchroSmsDto {
	return r._body
}
