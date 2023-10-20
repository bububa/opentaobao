package aliqin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinFcDigitalsmsCreatetemplateAPIRequest 数字短信模板创建 API请求
// alibaba.aliqin.fc.digitalsms.createtemplate
//
// 数字短信模板创建
type AlibabaAliqinFcDigitalsmsCreatetemplateAPIRequest struct {
	model.Params
	// 系统自动生成
	_templateContents []DigitalSmsTemplateContentDto
	// 模板名称
	_templateName string
	// 申请说明
	_applyRemark string
}

// NewAlibabaAliqinFcDigitalsmsCreatetemplateRequest 初始化AlibabaAliqinFcDigitalsmsCreatetemplateAPIRequest对象
func NewAlibabaAliqinFcDigitalsmsCreatetemplateRequest() *AlibabaAliqinFcDigitalsmsCreatetemplateAPIRequest {
	return &AlibabaAliqinFcDigitalsmsCreatetemplateAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAliqinFcDigitalsmsCreatetemplateAPIRequest) Reset() {
	r._templateContents = r._templateContents[:0]
	r._templateName = ""
	r._applyRemark = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAliqinFcDigitalsmsCreatetemplateAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.fc.digitalsms.createtemplate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAliqinFcDigitalsmsCreatetemplateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAliqinFcDigitalsmsCreatetemplateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTemplateContents is TemplateContents Setter
// 系统自动生成
func (r *AlibabaAliqinFcDigitalsmsCreatetemplateAPIRequest) SetTemplateContents(_templateContents []DigitalSmsTemplateContentDto) error {
	r._templateContents = _templateContents
	r.Set("template_contents", _templateContents)
	return nil
}

// GetTemplateContents TemplateContents Getter
func (r AlibabaAliqinFcDigitalsmsCreatetemplateAPIRequest) GetTemplateContents() []DigitalSmsTemplateContentDto {
	return r._templateContents
}

// SetTemplateName is TemplateName Setter
// 模板名称
func (r *AlibabaAliqinFcDigitalsmsCreatetemplateAPIRequest) SetTemplateName(_templateName string) error {
	r._templateName = _templateName
	r.Set("template_name", _templateName)
	return nil
}

// GetTemplateName TemplateName Getter
func (r AlibabaAliqinFcDigitalsmsCreatetemplateAPIRequest) GetTemplateName() string {
	return r._templateName
}

// SetApplyRemark is ApplyRemark Setter
// 申请说明
func (r *AlibabaAliqinFcDigitalsmsCreatetemplateAPIRequest) SetApplyRemark(_applyRemark string) error {
	r._applyRemark = _applyRemark
	r.Set("apply_remark", _applyRemark)
	return nil
}

// GetApplyRemark ApplyRemark Getter
func (r AlibabaAliqinFcDigitalsmsCreatetemplateAPIRequest) GetApplyRemark() string {
	return r._applyRemark
}

var poolAlibabaAliqinFcDigitalsmsCreatetemplateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAliqinFcDigitalsmsCreatetemplateRequest()
	},
}

// GetAlibabaAliqinFcDigitalsmsCreatetemplateRequest 从 sync.Pool 获取 AlibabaAliqinFcDigitalsmsCreatetemplateAPIRequest
func GetAlibabaAliqinFcDigitalsmsCreatetemplateAPIRequest() *AlibabaAliqinFcDigitalsmsCreatetemplateAPIRequest {
	return poolAlibabaAliqinFcDigitalsmsCreatetemplateAPIRequest.Get().(*AlibabaAliqinFcDigitalsmsCreatetemplateAPIRequest)
}

// ReleaseAlibabaAliqinFcDigitalsmsCreatetemplateAPIRequest 将 AlibabaAliqinFcDigitalsmsCreatetemplateAPIRequest 放入 sync.Pool
func ReleaseAlibabaAliqinFcDigitalsmsCreatetemplateAPIRequest(v *AlibabaAliqinFcDigitalsmsCreatetemplateAPIRequest) {
	v.Reset()
	poolAlibabaAliqinFcDigitalsmsCreatetemplateAPIRequest.Put(v)
}
