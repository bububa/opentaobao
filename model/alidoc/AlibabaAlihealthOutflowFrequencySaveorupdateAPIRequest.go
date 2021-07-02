package alidoc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthOutflowFrequencySaveorupdateAPIRequest 处方外流-药品频次同步接口 API请求
// alibaba.alihealth.outflow.frequency.saveorupdate
//
// 处方外流-药品频次同步接口
type AlibabaAlihealthOutflowFrequencySaveorupdateAPIRequest struct {
	model.Params
	// 系统自动生成
	_frequencyRequest *FrequencyRequest
}

// NewAlibabaAlihealthOutflowFrequencySaveorupdateRequest 初始化AlibabaAlihealthOutflowFrequencySaveorupdateAPIRequest对象
func NewAlibabaAlihealthOutflowFrequencySaveorupdateRequest() *AlibabaAlihealthOutflowFrequencySaveorupdateAPIRequest {
	return &AlibabaAlihealthOutflowFrequencySaveorupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthOutflowFrequencySaveorupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.outflow.frequency.saveorupdate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthOutflowFrequencySaveorupdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetFrequencyRequest is FrequencyRequest Setter
// 系统自动生成
func (r *AlibabaAlihealthOutflowFrequencySaveorupdateAPIRequest) SetFrequencyRequest(_frequencyRequest *FrequencyRequest) error {
	r._frequencyRequest = _frequencyRequest
	r.Set("frequency_request", _frequencyRequest)
	return nil
}

// GetFrequencyRequest FrequencyRequest Getter
func (r AlibabaAlihealthOutflowFrequencySaveorupdateAPIRequest) GetFrequencyRequest() *FrequencyRequest {
	return r._frequencyRequest
}
