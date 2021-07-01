package aliqin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIsvDigitalsmsCreatetemplateAPIRequest
数字短信模板创建 API请求
alibaba.isv.digitalsms.createtemplate

数字短信模板创建，给聚石塔，类型：2 */
type AlibabaIsvDigitalsmsCreatetemplateAPIRequest struct {
	model.Params
	// 模板名称
	_templateName string
	// 系统自动生成
	_templateContents []DigitalSmsTemplateContentDto
	// 申请说明
	_applyRemark string
}

// New
