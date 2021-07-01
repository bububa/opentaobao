package alidoc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthOutflowFrequencySaveorupdateAPIRequest
处方外流-药品频次同步接口 API请求
alibaba.alihealth.outflow.frequency.saveorupdate

处方外流-药品频次同步接口 */
type AlibabaAlihealthOutflowFrequencySaveorupdateAPIRequest struct {
	model.Params
	// 系统自动生成
	_frequencyRequest *FrequencyRequest
}

// New
