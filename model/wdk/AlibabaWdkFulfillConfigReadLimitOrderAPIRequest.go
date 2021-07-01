package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkFulfillConfigReadLimitOrderAPIRequest
根据仓code查询仓限单配置 API请求
alibaba.wdk.fulfill.config.read.limit.order

根据仓code查询仓限单配置 */
type AlibabaWdkFulfillConfigReadLimitOrderAPIRequest struct {
	model.Params
	// 仓code集合
	_warehouseCodeList []string
}

// NewAlibabaWdkFulfillConfigReadLimitOrderRequest 初始化AlibabaWdkFulfillConfigReadLimitOrderAPIRequest对象
func NewAlibabaWdkFulfillConfigReadLimitOrderRequest() *AlibabaWdkFulfillConfigReadLimitOrderAPIRequest {
	return &AlibabaWdkFulfillConfigReadLimitOrderAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkFulfillConfigReadLimitOrderAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.fulfill.config.read.limit.order"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkFulfillConfigReadLimitOrderAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is WarehouseCodeList Setter
// 仓code集合
func (r *AlibabaWdkFulfillConfigReadLimitOrderAPIRequest) SetWarehouseCodeList(_warehouseCodeList []string) error {
	r._warehouseCodeList = _warehouseCodeList
	r.Set("warehouse_code_list", _warehouseCodeList)
	return nil
}

// Get WarehouseCodeList Getter
func (r AlibabaWdkFulfillConfigReadLimitOrderAPIRequest) GetWarehouseCodeList() []string {
	return r._warehouseCodeList
}
