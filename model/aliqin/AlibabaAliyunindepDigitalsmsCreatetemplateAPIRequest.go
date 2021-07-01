package aliqin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
数字短信模板创建 API请求
alibaba.aliyunindep.digitalsms.createtemplate

数字短信模板创建，给阿里云一方产品使用，类型：9
*/
type AlibabaAliyunindepDigitalsmsCreatetemplateAPIRequest struct {
    model.Params
    // 模板名称
    _templateName   string
    // 系统自动生成
    _templateContents   []DigitalSmsTemplateContentDto
    // 申请说明
    _applyRemark   string
}

// 初始化AlibabaAliyunindepDigitalsmsCreatetemplateAPIRequest对象
func NewAlibabaAliyunindepDigitalsmsCreatetemplateRequest() *AlibabaAliyunindepDigitalsmsCreatetemplateAPIRequest{
    return &AlibabaAliyunindepDigitalsmsCreatetemplateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAliyunindepDigitalsmsCreatetemplateAPIRequest) GetApiMethodName() string {
    return "alibaba.aliyunindep.digitalsms.createtemplate"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAliyunindepDigitalsmsCreatetemplateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TemplateName Setter
// 模板名称
func (r *AlibabaAliyunindepDigitalsmsCreatetemplateAPIRequest) SetTemplateName(_templateName string) error {
    r._templateName = _templateName
    r.Set("template_name", _templateName)
    return nil
}

// TemplateName Getter
func (r AlibabaAliyunindepDigitalsmsCreatetemplateAPIRequest) GetTemplateName() string {
    return r._templateName
}
// TemplateContents Setter
// 系统自动生成
func (r *AlibabaAliyunindepDigitalsmsCreatetemplateAPIRequest) SetTemplateContents(_templateContents []DigitalSmsTemplateContentDto) error {
    r._templateContents = _templateContents
    r.Set("template_contents", _templateContents)
    return nil
}

// TemplateContents Getter
func (r AlibabaAliyunindepDigitalsmsCreatetemplateAPIRequest) GetTemplateContents() []DigitalSmsTemplateContentDto {
    return r._templateContents
}
// ApplyRemark Setter
// 申请说明
func (r *AlibabaAliyunindepDigitalsmsCreatetemplateAPIRequest) SetApplyRemark(_applyRemark string) error {
    r._applyRemark = _applyRemark
    r.Set("apply_remark", _applyRemark)
    return nil
}

// ApplyRemark Getter
func (r AlibabaAliyunindepDigitalsmsCreatetemplateAPIRequest) GetApplyRemark() string {
    return r._applyRemark
}
