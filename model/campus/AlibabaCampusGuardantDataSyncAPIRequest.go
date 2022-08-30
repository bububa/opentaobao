package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusGuardantDataSyncAPIRequest 刷卡数据同步 API请求
// alibaba.campus.guardant.data.sync
//
// 数据同步门禁系统
type AlibabaCampusGuardantDataSyncAPIRequest struct {
	model.Params
	// 1-刷卡流水
	_dataType string
	// 供应商名称
	_supplierName string
	// json串
	_data string
}

// NewAlibabaCampusGuardantDataSyncRequest 初始化AlibabaCampusGuardantDataSyncAPIRequest对象
func NewAlibabaCampusGuardantDataSyncRequest() *AlibabaCampusGuardantDataSyncAPIRequest {
	return &AlibabaCampusGuardantDataSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusGuardantDataSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.guardant.data.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusGuardantDataSyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetDataType is DataType Setter
// 1-刷卡流水
func (r *AlibabaCampusGuardantDataSyncAPIRequest) SetDataType(_dataType string) error {
	r._dataType = _dataType
	r.Set("data_type", _dataType)
	return nil
}

// GetDataType DataType Getter
func (r AlibabaCampusGuardantDataSyncAPIRequest) GetDataType() string {
	return r._dataType
}

// SetSupplierName is SupplierName Setter
// 供应商名称
func (r *AlibabaCampusGuardantDataSyncAPIRequest) SetSupplierName(_supplierName string) error {
	r._supplierName = _supplierName
	r.Set("supplier_name", _supplierName)
	return nil
}

// GetSupplierName SupplierName Getter
func (r AlibabaCampusGuardantDataSyncAPIRequest) GetSupplierName() string {
	return r._supplierName
}

// SetData is Data Setter
// json串
func (r *AlibabaCampusGuardantDataSyncAPIRequest) SetData(_data string) error {
	r._data = _data
	r.Set("data", _data)
	return nil
}

// GetData Data Getter
func (r AlibabaCampusGuardantDataSyncAPIRequest) GetData() string {
	return r._data
}
