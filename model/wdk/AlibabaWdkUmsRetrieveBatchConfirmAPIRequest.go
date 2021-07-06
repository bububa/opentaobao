package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkUmsRetrieveBatchConfirmAPIRequest 批量消息确认 API请求
// alibaba.wdk.ums.retrieve.batch.confirm
//
// 批量消息确认
type AlibabaWdkUmsRetrieveBatchConfirmAPIRequest struct {
	model.Params
	// warehouse_code
	_uuids []string
	// warehouse_code
	_warehouseCode string
}

// NewAlibabaWdkUmsRetrieveBatchConfirmRequest 初始化AlibabaWdkUmsRetrieveBatchConfirmAPIRequest对象
func NewAlibabaWdkUmsRetrieveBatchConfirmRequest() *AlibabaWdkUmsRetrieveBatchConfirmAPIRequest {
	return &AlibabaWdkUmsRetrieveBatchConfirmAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkUmsRetrieveBatchConfirmAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.ums.retrieve.batch.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkUmsRetrieveBatchConfirmAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetUuids is Uuids Setter
// warehouse_code
func (r *AlibabaWdkUmsRetrieveBatchConfirmAPIRequest) SetUuids(_uuids []string) error {
	r._uuids = _uuids
	r.Set("uuids", _uuids)
	return nil
}

// GetUuids Uuids Getter
func (r AlibabaWdkUmsRetrieveBatchConfirmAPIRequest) GetUuids() []string {
	return r._uuids
}

// SetWarehouseCode is WarehouseCode Setter
// warehouse_code
func (r *AlibabaWdkUmsRetrieveBatchConfirmAPIRequest) SetWarehouseCode(_warehouseCode string) error {
	r._warehouseCode = _warehouseCode
	r.Set("warehouse_code", _warehouseCode)
	return nil
}

// GetWarehouseCode WarehouseCode Getter
func (r AlibabaWdkUmsRetrieveBatchConfirmAPIRequest) GetWarehouseCode() string {
	return r._warehouseCode
}
