package alihealthoutflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthOutflowDrugSaveorupdateAPIRequest 处方外流-药品同步接口 API请求
// alibaba.alihealth.outflow.drug.saveorupdate
//
// 处方外流-药品同步接口
type AlibabaAlihealthOutflowDrugSaveorupdateAPIRequest struct {
	model.Params
	// 结果集
	_drugRequest *DrugRequest
}

// NewAlibabaAlihealthOutflowDrugSaveorupdateRequest 初始化AlibabaAlihealthOutflowDrugSaveorupdateAPIRequest对象
func NewAlibabaAlihealthOutflowDrugSaveorupdateRequest() *AlibabaAlihealthOutflowDrugSaveorupdateAPIRequest {
	return &AlibabaAlihealthOutflowDrugSaveorupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthOutflowDrugSaveorupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.outflow.drug.saveorupdate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthOutflowDrugSaveorupdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetDrugRequest is DrugRequest Setter
// 结果集
func (r *AlibabaAlihealthOutflowDrugSaveorupdateAPIRequest) SetDrugRequest(_drugRequest *DrugRequest) error {
	r._drugRequest = _drugRequest
	r.Set("drug_request", _drugRequest)
	return nil
}

// GetDrugRequest DrugRequest Getter
func (r AlibabaAlihealthOutflowDrugSaveorupdateAPIRequest) GetDrugRequest() *DrugRequest {
	return r._drugRequest
}
