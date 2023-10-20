package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkumsretrievebatchconfirmAPIRequest 批量消息确认 API请求
// alibaba.wdk.ums.retrieve.batch.confirm
//
// 批量消息确认
type AlibabawdkumsretrievebatchconfirmAPIRequest struct {
	model.Params
	// warehouse_code
	_uuids []string
	// warehouse_code
	_warehouseCode string
}

// NewAlibabawdkumsretrievebatchconfirmRequest 初始化AlibabawdkumsretrievebatchconfirmAPIRequest对象
func NewAlibabawdkumsretrievebatchconfirmRequest() *AlibabawdkumsretrievebatchconfirmAPIRequest {
	return &AlibabawdkumsretrievebatchconfirmAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkumsretrievebatchconfirmAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.ums.retrieve.batch.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkumsretrievebatchconfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkumsretrievebatchconfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUuids is Uuids Setter
// warehouse_code
func (r *AlibabawdkumsretrievebatchconfirmAPIRequest) SetUuids(_uuids []string) error {
	r._uuids = _uuids
	r.Set("uuids", _uuids)
	return nil
}

// GetUuids Uuids Getter
func (r AlibabawdkumsretrievebatchconfirmAPIRequest) GetUuids() []string {
	return r._uuids
}

// SetWarehouseCode is WarehouseCode Setter
// warehouse_code
func (r *AlibabawdkumsretrievebatchconfirmAPIRequest) SetWarehouseCode(_warehouseCode string) error {
	r._warehouseCode = _warehouseCode
	r.Set("warehouse_code", _warehouseCode)
	return nil
}

// GetWarehouseCode WarehouseCode Getter
func (r AlibabawdkumsretrievebatchconfirmAPIRequest) GetWarehouseCode() string {
	return r._warehouseCode
}
