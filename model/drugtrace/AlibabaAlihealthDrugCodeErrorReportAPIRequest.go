package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugCodeErrorReportAPIRequest
码信息错误上报 API请求
alibaba.alihealth.drug.code.error.report

提供码信息错误上报功能，用于数据校对 */
type AlibabaAlihealthDrugCodeErrorReportAPIRequest struct {
	model.Params
	// 追溯码
	_code string
	// 有问题的字段名称
	_fieldName string
	// 通过码获得的问题字段值
	_codeValue string
	// 平台获得/期望的问题字段值
	_sourceValue string
	// 错误信息描述
	_errMsg string
	// 上报人员
	_reporter string
	// 上报人员邮箱
	_reporterEmail string
	// 上报人员手机号
	_reporterMobile string
}

// New
