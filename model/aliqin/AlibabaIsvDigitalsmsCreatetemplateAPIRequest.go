package aliqin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIsvDigitalsmsCreatetemplateAPIRequest 数字短信模板创建 API请求
// alibaba.isv.digitalsms.createtemplate
//
// 数字短信模板创建，给聚石塔，类型：2
type AlibabaIsvDigitalsmsCreatetemplateAPIRequest struct {
	model.Params
	// 系统自动生成
	_templateContents []DigitalSmsTemplateContentDto
	// 模板名称
	_templateName string
	// 申请说明
	_applyRemark string
}

// NewAlibabaIsvDigitalsmsCreatetemplateRequest 初始化AlibabaIsvDigitalsmsCreatetemplateAPIRequest对象
func NewAlibabaIsvDigitalsmsCreatetemplateRequest() *AlibabaIsvDigitalsmsCreatetemplateAPIRequest {
	return &AlibabaIsvDigitalsmsCreatetemplateAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIsvDigitalsmsCreatetemplateAPIRequest) Reset() {
	r._templateContents = r._templateContents[:0]
	r._templateName = ""
	r._applyRemark = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIsvDigitalsmsCreatetemplateAPIRequest) GetApiMethodName() string {
	return "alibaba.isv.digitalsms.createtemplate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIsvDigitalsmsCreatetemplateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIsvDigitalsmsCreatetemplateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTemplateContents is TemplateContents Setter
// 系统自动生成
func (r *AlibabaIsvDigitalsmsCreatetemplateAPIRequest) SetTemplateContents(_templateContents []DigitalSmsTemplateContentDto) error {
	r._templateContents = _templateContents
	r.Set("template_contents", _templateContents)
	return nil
}

// GetTemplateContents TemplateContents Getter
func (r AlibabaIsvDigitalsmsCreatetemplateAPIRequest) GetTemplateContents() []DigitalSmsTemplateContentDto {
	return r._templateContents
}

// SetTemplateName is TemplateName Setter
// 模板名称
func (r *AlibabaIsvDigitalsmsCreatetemplateAPIRequest) SetTemplateName(_templateName string) error {
	r._templateName = _templateName
	r.Set("template_name", _templateName)
	return nil
}

// GetTemplateName TemplateName Getter
func (r AlibabaIsvDigitalsmsCreatetemplateAPIRequest) GetTemplateName() string {
	return r._templateName
}

// SetApplyRemark is ApplyRemark Setter
// 申请说明
func (r *AlibabaIsvDigitalsmsCreatetemplateAPIRequest) SetApplyRemark(_applyRemark string) error {
	r._applyRemark = _applyRemark
	r.Set("apply_remark", _applyRemark)
	return nil
}

// GetApplyRemark ApplyRemark Getter
func (r AlibabaIsvDigitalsmsCreatetemplateAPIRequest) GetApplyRemark() string {
	return r._applyRemark
}

var poolAlibabaIsvDigitalsmsCreatetemplateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIsvDigitalsmsCreatetemplateRequest()
	},
}

// GetAlibabaIsvDigitalsmsCreatetemplateRequest 从 sync.Pool 获取 AlibabaIsvDigitalsmsCreatetemplateAPIRequest
func GetAlibabaIsvDigitalsmsCreatetemplateAPIRequest() *AlibabaIsvDigitalsmsCreatetemplateAPIRequest {
	return poolAlibabaIsvDigitalsmsCreatetemplateAPIRequest.Get().(*AlibabaIsvDigitalsmsCreatetemplateAPIRequest)
}

// ReleaseAlibabaIsvDigitalsmsCreatetemplateAPIRequest 将 AlibabaIsvDigitalsmsCreatetemplateAPIRequest 放入 sync.Pool
func ReleaseAlibabaIsvDigitalsmsCreatetemplateAPIRequest(v *AlibabaIsvDigitalsmsCreatetemplateAPIRequest) {
	v.Reset()
	poolAlibabaIsvDigitalsmsCreatetemplateAPIRequest.Put(v)
}
