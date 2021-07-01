package alihealthoutflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthOutflowDrugSaveorupdateAPIRequest
处方外流-药品同步接口 API请求
alibaba.alihealth.outflow.drug.saveorupdate

处方外流-药品同步接口 */
type AlibabaAlihealthOutflowDrugSaveorupdateAPIRequest struct {
	model.Params
	// 结果集
	_drugRequest *DrugRequest
}

// New
