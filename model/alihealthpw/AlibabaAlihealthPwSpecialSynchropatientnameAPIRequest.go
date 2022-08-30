package alihealthpw

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthPwSpecialSynchropatientnameAPIRequest 同步患者姓名至阿里健康 API请求
// alibaba.alihealth.pw.special.synchropatientname
//
// 同步患者姓名至阿里健康
type AlibabaAlihealthPwSpecialSynchropatientnameAPIRequest struct {
	model.Params
	// 入参
	_body *SynchroPatientNameDto
}

// NewAlibabaAlihealthPwSpecialSynchropatientnameRequest 初始化AlibabaAlihealthPwSpecialSynchropatientnameAPIRequest对象
func NewAlibabaAlihealthPwSpecialSynchropatientnameRequest() *AlibabaAlihealthPwSpecialSynchropatientnameAPIRequest {
	return &AlibabaAlihealthPwSpecialSynchropatientnameAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthPwSpecialSynchropatientnameAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.pw.special.synchropatientname"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthPwSpecialSynchropatientnameAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetBody is Body Setter
// 入参
func (r *AlibabaAlihealthPwSpecialSynchropatientnameAPIRequest) SetBody(_body *SynchroPatientNameDto) error {
	r._body = _body
	r.Set("body", _body)
	return nil
}

// GetBody Body Getter
func (r AlibabaAlihealthPwSpecialSynchropatientnameAPIRequest) GetBody() *SynchroPatientNameDto {
	return r._body
}
