package aliqin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaisvdigitalsmscreatetemplateAPIRequest 数字短信模板创建 API请求
// alibaba.isv.digitalsms.createtemplate
//
// 数字短信模板创建，给聚石塔，类型：2
type AlibabaisvdigitalsmscreatetemplateAPIRequest struct {
	model.Params
	// 系统自动生成
	_templateContents []DigitalSmsTemplateContentDto
	// 模板名称
	_templateName string
	// 申请说明
	_applyRemark string
}

// NewAlibabaisvdigitalsmscreatetemplateRequest 初始化AlibabaisvdigitalsmscreatetemplateAPIRequest对象
func NewAlibabaisvdigitalsmscreatetemplateRequest() *AlibabaisvdigitalsmscreatetemplateAPIRequest {
	return &AlibabaisvdigitalsmscreatetemplateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaisvdigitalsmscreatetemplateAPIRequest) GetApiMethodName() string {
	return "alibaba.isv.digitalsms.createtemplate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaisvdigitalsmscreatetemplateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaisvdigitalsmscreatetemplateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTemplateContents is TemplateContents Setter
// 系统自动生成
func (r *AlibabaisvdigitalsmscreatetemplateAPIRequest) SetTemplateContents(_templateContents []DigitalSmsTemplateContentDto) error {
	r._templateContents = _templateContents
	r.Set("template_contents", _templateContents)
	return nil
}

// GetTemplateContents TemplateContents Getter
func (r AlibabaisvdigitalsmscreatetemplateAPIRequest) GetTemplateContents() []DigitalSmsTemplateContentDto {
	return r._templateContents
}

// SetTemplateName is TemplateName Setter
// 模板名称
func (r *AlibabaisvdigitalsmscreatetemplateAPIRequest) SetTemplateName(_templateName string) error {
	r._templateName = _templateName
	r.Set("template_name", _templateName)
	return nil
}

// GetTemplateName TemplateName Getter
func (r AlibabaisvdigitalsmscreatetemplateAPIRequest) GetTemplateName() string {
	return r._templateName
}

// SetApplyRemark is ApplyRemark Setter
// 申请说明
func (r *AlibabaisvdigitalsmscreatetemplateAPIRequest) SetApplyRemark(_applyRemark string) error {
	r._applyRemark = _applyRemark
	r.Set("apply_remark", _applyRemark)
	return nil
}

// GetApplyRemark ApplyRemark Getter
func (r AlibabaisvdigitalsmscreatetemplateAPIRequest) GetApplyRemark() string {
	return r._applyRemark
}
