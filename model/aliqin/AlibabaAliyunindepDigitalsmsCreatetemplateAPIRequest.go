package aliqin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliyunindepDigitalsmsCreatetemplateAPIRequest 数字短信模板创建 API请求
// alibaba.aliyunindep.digitalsms.createtemplate
//
// 数字短信模板创建，给阿里云一方产品使用，类型：9
type AlibabaAliyunindepDigitalsmsCreatetemplateAPIRequest struct {
	model.Params
	// 系统自动生成
	_templateContents []DigitalSmsTemplateContentDto
	// 模板名称
	_templateName string
	// 申请说明
	_applyRemark string
}

// NewAlibabaAliyunindepDigitalsmsCreatetemplateRequest 初始化AlibabaAliyunindepDigitalsmsCreatetemplateAPIRequest对象
func NewAlibabaAliyunindepDigitalsmsCreatetemplateRequest() *AlibabaAliyunindepDigitalsmsCreatetemplateAPIRequest {
	return &AlibabaAliyunindepDigitalsmsCreatetemplateAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAliyunindepDigitalsmsCreatetemplateAPIRequest) Reset() {
	r._templateContents = r._templateContents[:0]
	r._templateName = ""
	r._applyRemark = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAliyunindepDigitalsmsCreatetemplateAPIRequest) GetApiMethodName() string {
	return "alibaba.aliyunindep.digitalsms.createtemplate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAliyunindepDigitalsmsCreatetemplateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAliyunindepDigitalsmsCreatetemplateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTemplateContents is TemplateContents Setter
// 系统自动生成
func (r *AlibabaAliyunindepDigitalsmsCreatetemplateAPIRequest) SetTemplateContents(_templateContents []DigitalSmsTemplateContentDto) error {
	r._templateContents = _templateContents
	r.Set("template_contents", _templateContents)
	return nil
}

// GetTemplateContents TemplateContents Getter
func (r AlibabaAliyunindepDigitalsmsCreatetemplateAPIRequest) GetTemplateContents() []DigitalSmsTemplateContentDto {
	return r._templateContents
}

// SetTemplateName is TemplateName Setter
// 模板名称
func (r *AlibabaAliyunindepDigitalsmsCreatetemplateAPIRequest) SetTemplateName(_templateName string) error {
	r._templateName = _templateName
	r.Set("template_name", _templateName)
	return nil
}

// GetTemplateName TemplateName Getter
func (r AlibabaAliyunindepDigitalsmsCreatetemplateAPIRequest) GetTemplateName() string {
	return r._templateName
}

// SetApplyRemark is ApplyRemark Setter
// 申请说明
func (r *AlibabaAliyunindepDigitalsmsCreatetemplateAPIRequest) SetApplyRemark(_applyRemark string) error {
	r._applyRemark = _applyRemark
	r.Set("apply_remark", _applyRemark)
	return nil
}

// GetApplyRemark ApplyRemark Getter
func (r AlibabaAliyunindepDigitalsmsCreatetemplateAPIRequest) GetApplyRemark() string {
	return r._applyRemark
}

var poolAlibabaAliyunindepDigitalsmsCreatetemplateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAliyunindepDigitalsmsCreatetemplateRequest()
	},
}

// GetAlibabaAliyunindepDigitalsmsCreatetemplateRequest 从 sync.Pool 获取 AlibabaAliyunindepDigitalsmsCreatetemplateAPIRequest
func GetAlibabaAliyunindepDigitalsmsCreatetemplateAPIRequest() *AlibabaAliyunindepDigitalsmsCreatetemplateAPIRequest {
	return poolAlibabaAliyunindepDigitalsmsCreatetemplateAPIRequest.Get().(*AlibabaAliyunindepDigitalsmsCreatetemplateAPIRequest)
}

// ReleaseAlibabaAliyunindepDigitalsmsCreatetemplateAPIRequest 将 AlibabaAliyunindepDigitalsmsCreatetemplateAPIRequest 放入 sync.Pool
func ReleaseAlibabaAliyunindepDigitalsmsCreatetemplateAPIRequest(v *AlibabaAliyunindepDigitalsmsCreatetemplateAPIRequest) {
	v.Reset()
	poolAlibabaAliyunindepDigitalsmsCreatetemplateAPIRequest.Put(v)
}
