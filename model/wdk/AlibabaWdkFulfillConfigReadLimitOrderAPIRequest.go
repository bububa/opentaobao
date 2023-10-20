package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkfulfillconfigreadlimitorderAPIRequest 根据仓code查询仓限单配置 API请求
// alibaba.wdk.fulfill.config.read.limit.order
//
// 根据仓code查询仓限单配置
type AlibabawdkfulfillconfigreadlimitorderAPIRequest struct {
	model.Params
	// 仓code集合
	_warehouseCodeList []string
}

// NewAlibabawdkfulfillconfigreadlimitorderRequest 初始化AlibabawdkfulfillconfigreadlimitorderAPIRequest对象
func NewAlibabawdkfulfillconfigreadlimitorderRequest() *AlibabawdkfulfillconfigreadlimitorderAPIRequest {
	return &AlibabawdkfulfillconfigreadlimitorderAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkfulfillconfigreadlimitorderAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.fulfill.config.read.limit.order"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkfulfillconfigreadlimitorderAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkfulfillconfigreadlimitorderAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWarehouseCodeList is WarehouseCodeList Setter
// 仓code集合
func (r *AlibabawdkfulfillconfigreadlimitorderAPIRequest) SetWarehouseCodeList(_warehouseCodeList []string) error {
	r._warehouseCodeList = _warehouseCodeList
	r.Set("warehouse_code_list", _warehouseCodeList)
	return nil
}

// GetWarehouseCodeList WarehouseCodeList Getter
func (r AlibabawdkfulfillconfigreadlimitorderAPIRequest) GetWarehouseCodeList() []string {
	return r._warehouseCodeList
}
