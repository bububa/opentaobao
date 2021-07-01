package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkUmsRetrieveConfirmAPIRequest
回流单－外部对已拉取到的UMS单据进行确认 API请求
alibaba.wdk.ums.retrieve.confirm

回流单－外部对已拉取到的UMS单据进行确认 */
type AlibabaWdkUmsRetrieveConfirmAPIRequest struct {
	model.Params
	// 店仓code，指的是作业对象，对应一个物理店或仓编码
	_warehouseCode string
	// 唯一识别码
	_uuid string
}

// NewAlibabaWdkUmsRetrieveConfirmRequest 初始化AlibabaWdkUmsRetrieveConfirmAPIRequest对象
func NewAlibabaWdkUmsRetrieveConfirmRequest() *AlibabaWdkUmsRetrieveConfirmAPIRequest {
	return &AlibabaWdkUmsRetrieveConfirmAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkUmsRetrieveConfirmAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.ums.retrieve.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkUmsRetrieveConfirmAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is WarehouseCode Setter
// 店仓code，指的是作业对象，对应一个物理店或仓编码
func (r *AlibabaWdkUmsRetrieveConfirmAPIRequest) SetWarehouseCode(_warehouseCode string) error {
	r._warehouseCode = _warehouseCode
	r.Set("warehouse_code", _warehouseCode)
	return nil
}

// Get WarehouseCode Getter
func (r AlibabaWdkUmsRetrieveConfirmAPIRequest) GetWarehouseCode() string {
	return r._warehouseCode
}

// Set is Uuid Setter
// 唯一识别码
func (r *AlibabaWdkUmsRetrieveConfirmAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// Get Uuid Getter
func (r AlibabaWdkUmsRetrieveConfirmAPIRequest) GetUuid() string {
	return r._uuid
}
