package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusguardantdatasyncAPIRequest 刷卡数据同步 API请求
// alibaba.campus.guardant.data.sync
//
// 数据同步门禁系统
type AlibabacampusguardantdatasyncAPIRequest struct {
	model.Params
	// 1-刷卡流水
	_dataType string
	// 供应商名称
	_supplierName string
	// json串
	_data string
}

// NewAlibabacampusguardantdatasyncRequest 初始化AlibabacampusguardantdatasyncAPIRequest对象
func NewAlibabacampusguardantdatasyncRequest() *AlibabacampusguardantdatasyncAPIRequest {
	return &AlibabacampusguardantdatasyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacampusguardantdatasyncAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.guardant.data.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacampusguardantdatasyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacampusguardantdatasyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDataType is DataType Setter
// 1-刷卡流水
func (r *AlibabacampusguardantdatasyncAPIRequest) SetDataType(_dataType string) error {
	r._dataType = _dataType
	r.Set("data_type", _dataType)
	return nil
}

// GetDataType DataType Getter
func (r AlibabacampusguardantdatasyncAPIRequest) GetDataType() string {
	return r._dataType
}

// SetSupplierName is SupplierName Setter
// 供应商名称
func (r *AlibabacampusguardantdatasyncAPIRequest) SetSupplierName(_supplierName string) error {
	r._supplierName = _supplierName
	r.Set("supplier_name", _supplierName)
	return nil
}

// GetSupplierName SupplierName Getter
func (r AlibabacampusguardantdatasyncAPIRequest) GetSupplierName() string {
	return r._supplierName
}

// SetData is Data Setter
// json串
func (r *AlibabacampusguardantdatasyncAPIRequest) SetData(_data string) error {
	r._data = _data
	r.Set("data", _data)
	return nil
}

// GetData Data Getter
func (r AlibabacampusguardantdatasyncAPIRequest) GetData() string {
	return r._data
}
