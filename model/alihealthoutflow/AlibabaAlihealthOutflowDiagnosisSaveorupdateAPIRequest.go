package alihealthoutflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthOutflowDiagnosisSaveorupdateAPIRequest
处方外流-诊断字典表 API请求
alibaba.alihealth.outflow.diagnosis.saveorupdate

阿里健康-处方外流-对外提供诊断字典表维护功能 */
type AlibabaAlihealthOutflowDiagnosisSaveorupdateAPIRequest struct {
	model.Params
	// 诊断数据
	_diagnoseDictDto *DiagnoseDictDto
}

// New
