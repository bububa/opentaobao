package aliqin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAliyunindepDigitalsmsCreatetemplateAPIRequest
数字短信模板创建 API请求
alibaba.aliyunindep.digitalsms.createtemplate

数字短信模板创建，给阿里云一方产品使用，类型：9 */
type AlibabaAliyunindepDigitalsmsCreatetemplateAPIRequest struct {
	model.Params
	// 模板名称
	_templateName string
	// 系统自动生成
	_templateContents []DigitalSmsTemplateContentDto
	// 申请说明
	_applyRemark string
}

// New
