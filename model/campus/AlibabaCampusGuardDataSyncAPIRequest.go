package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusguarddatasyncAPIRequest 卡巴数据同步 API请求
// alibaba.campus.guard.data.sync
//
// 数据同步门禁系统
type AlibabacampusguarddatasyncAPIRequest struct {
	model.Params
	// 1-刷卡流水
	_dataType string
	// 供应商名称
	_supplierName string
	// json串
	_data string
}

// NewAlibabacampusguarddatasyncRequest 初始化AlibabacampusguarddatasyncAPIRequest对象
func NewAlibabacampusguarddatasyncRequest() *AlibabacampusguarddatasyncAPIRequest {
	return &AlibabacampusguarddatasyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacampusguarddatasyncAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.guard.data.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacampusguarddatasyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacampusguarddatasyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDataType is DataType Setter
// 1-刷卡流水
func (r *AlibabacampusguarddatasyncAPIRequest) SetDataType(_dataType string) error {
	r._dataType = _dataType
	r.Set("data_type", _dataType)
	return nil
}

// GetDataType DataType Getter
func (r AlibabacampusguarddatasyncAPIRequest) GetDataType() string {
	return r._dataType
}

// SetSupplierName is SupplierName Setter
// 供应商名称
func (r *AlibabacampusguarddatasyncAPIRequest) SetSupplierName(_supplierName string) error {
	r._supplierName = _supplierName
	r.Set("supplier_name", _supplierName)
	return nil
}

// GetSupplierName SupplierName Getter
func (r AlibabacampusguarddatasyncAPIRequest) GetSupplierName() string {
	return r._supplierName
}

// SetData is Data Setter
// json串
func (r *AlibabacampusguarddatasyncAPIRequest) SetData(_data string) error {
	r._data = _data
	r.Set("data", _data)
	return nil
}

// GetData Data Getter
func (r AlibabacampusguarddatasyncAPIRequest) GetData() string {
	return r._data
}
