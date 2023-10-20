package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkFulfillConfigReadLimitOrderAPIRequest 根据仓code查询仓限单配置 API请求
// alibaba.wdk.fulfill.config.read.limit.order
//
// 根据仓code查询仓限单配置
type AlibabaWdkFulfillConfigReadLimitOrderAPIRequest struct {
	model.Params
	// 仓code集合
	_warehouseCodeList []string
}

// NewAlibabaWdkFulfillConfigReadLimitOrderRequest 初始化AlibabaWdkFulfillConfigReadLimitOrderAPIRequest对象
func NewAlibabaWdkFulfillConfigReadLimitOrderRequest() *AlibabaWdkFulfillConfigReadLimitOrderAPIRequest {
	return &AlibabaWdkFulfillConfigReadLimitOrderAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkFulfillConfigReadLimitOrderAPIRequest) Reset() {
	r._warehouseCodeList = r._warehouseCodeList[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkFulfillConfigReadLimitOrderAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.fulfill.config.read.limit.order"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkFulfillConfigReadLimitOrderAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkFulfillConfigReadLimitOrderAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWarehouseCodeList is WarehouseCodeList Setter
// 仓code集合
func (r *AlibabaWdkFulfillConfigReadLimitOrderAPIRequest) SetWarehouseCodeList(_warehouseCodeList []string) error {
	r._warehouseCodeList = _warehouseCodeList
	r.Set("warehouse_code_list", _warehouseCodeList)
	return nil
}

// GetWarehouseCodeList WarehouseCodeList Getter
func (r AlibabaWdkFulfillConfigReadLimitOrderAPIRequest) GetWarehouseCodeList() []string {
	return r._warehouseCodeList
}

var poolAlibabaWdkFulfillConfigReadLimitOrderAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkFulfillConfigReadLimitOrderRequest()
	},
}

// GetAlibabaWdkFulfillConfigReadLimitOrderRequest 从 sync.Pool 获取 AlibabaWdkFulfillConfigReadLimitOrderAPIRequest
func GetAlibabaWdkFulfillConfigReadLimitOrderAPIRequest() *AlibabaWdkFulfillConfigReadLimitOrderAPIRequest {
	return poolAlibabaWdkFulfillConfigReadLimitOrderAPIRequest.Get().(*AlibabaWdkFulfillConfigReadLimitOrderAPIRequest)
}

// ReleaseAlibabaWdkFulfillConfigReadLimitOrderAPIRequest 将 AlibabaWdkFulfillConfigReadLimitOrderAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkFulfillConfigReadLimitOrderAPIRequest(v *AlibabaWdkFulfillConfigReadLimitOrderAPIRequest) {
	v.Reset()
	poolAlibabaWdkFulfillConfigReadLimitOrderAPIRequest.Put(v)
}
