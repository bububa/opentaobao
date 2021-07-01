package alihealthoutflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthOutflowUsageSaveorupdateAPIRequest
处方外流-用法字典表 API请求
alibaba.alihealth.outflow.usage.saveorupdate

阿里健康-处方外流-对外提供用法字典表维护功能 */
type AlibabaAlihealthOutflowUsageSaveorupdateAPIRequest struct {
	model.Params
	// 用法数据
	_usageRequest *UsageRequest
}

// New
