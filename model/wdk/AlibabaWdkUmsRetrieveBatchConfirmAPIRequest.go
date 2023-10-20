package wdk

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkUmsRetrieveBatchConfirmAPIRequest) Reset() {
	r._uuids = r._uuids[:0]
	r._warehouseCode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkUmsRetrieveBatchConfirmAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.ums.retrieve.batch.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkUmsRetrieveBatchConfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkUmsRetrieveBatchConfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaWdkUmsRetrieveBatchConfirmAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkUmsRetrieveBatchConfirmRequest()
	},
}

// GetAlibabaWdkUmsRetrieveBatchConfirmRequest 从 sync.Pool 获取 AlibabaWdkUmsRetrieveBatchConfirmAPIRequest
func GetAlibabaWdkUmsRetrieveBatchConfirmAPIRequest() *AlibabaWdkUmsRetrieveBatchConfirmAPIRequest {
	return poolAlibabaWdkUmsRetrieveBatchConfirmAPIRequest.Get().(*AlibabaWdkUmsRetrieveBatchConfirmAPIRequest)
}

// ReleaseAlibabaWdkUmsRetrieveBatchConfirmAPIRequest 将 AlibabaWdkUmsRetrieveBatchConfirmAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkUmsRetrieveBatchConfirmAPIRequest(v *AlibabaWdkUmsRetrieveBatchConfirmAPIRequest) {
	v.Reset()
	poolAlibabaWdkUmsRetrieveBatchConfirmAPIRequest.Put(v)
}
