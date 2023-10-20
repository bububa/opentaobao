package aliqin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaaliyunindepdigitalsmscreatetemplateAPIRequest 数字短信模板创建 API请求
// alibaba.aliyunindep.digitalsms.createtemplate
//
// 数字短信模板创建，给阿里云一方产品使用，类型：9
type AlibabaaliyunindepdigitalsmscreatetemplateAPIRequest struct {
	model.Params
	// 系统自动生成
	_templateContents []DigitalSmsTemplateContentDto
	// 模板名称
	_templateName string
	// 申请说明
	_applyRemark string
}

// NewAlibabaaliyunindepdigitalsmscreatetemplateRequest 初始化AlibabaaliyunindepdigitalsmscreatetemplateAPIRequest对象
func NewAlibabaaliyunindepdigitalsmscreatetemplateRequest() *AlibabaaliyunindepdigitalsmscreatetemplateAPIRequest {
	return &AlibabaaliyunindepdigitalsmscreatetemplateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaaliyunindepdigitalsmscreatetemplateAPIRequest) GetApiMethodName() string {
	return "alibaba.aliyunindep.digitalsms.createtemplate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaaliyunindepdigitalsmscreatetemplateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaaliyunindepdigitalsmscreatetemplateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTemplateContents is TemplateContents Setter
// 系统自动生成
func (r *AlibabaaliyunindepdigitalsmscreatetemplateAPIRequest) SetTemplateContents(_templateContents []DigitalSmsTemplateContentDto) error {
	r._templateContents = _templateContents
	r.Set("template_contents", _templateContents)
	return nil
}

// GetTemplateContents TemplateContents Getter
func (r AlibabaaliyunindepdigitalsmscreatetemplateAPIRequest) GetTemplateContents() []DigitalSmsTemplateContentDto {
	return r._templateContents
}

// SetTemplateName is TemplateName Setter
// 模板名称
func (r *AlibabaaliyunindepdigitalsmscreatetemplateAPIRequest) SetTemplateName(_templateName string) error {
	r._templateName = _templateName
	r.Set("template_name", _templateName)
	return nil
}

// GetTemplateName TemplateName Getter
func (r AlibabaaliyunindepdigitalsmscreatetemplateAPIRequest) GetTemplateName() string {
	return r._templateName
}

// SetApplyRemark is ApplyRemark Setter
// 申请说明
func (r *AlibabaaliyunindepdigitalsmscreatetemplateAPIRequest) SetApplyRemark(_applyRemark string) error {
	r._applyRemark = _applyRemark
	r.Set("apply_remark", _applyRemark)
	return nil
}

// GetApplyRemark ApplyRemark Getter
func (r AlibabaaliyunindepdigitalsmscreatetemplateAPIRequest) GetApplyRemark() string {
	return r._applyRemark
}
