package ascpchannel

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpChannelDistributorInventoryListGetAPIRequest 批量查询渠道库存 API请求
// alibaba.ascp.channel.distributor.inventory.list.get
//
// 淘外分销批量查询渠道产品的库存
type AlibabaAscpChannelDistributorInventoryListGetAPIRequest struct {
	model.Params
	// 系统自动生成
	_inventoryRequest *BatchChannelInventoryQuery
}

// NewAlibabaAscpChannelDistributorInventoryListGetRequest 初始化AlibabaAscpChannelDistributorInventoryListGetAPIRequest对象
func NewAlibabaAscpChannelDistributorInventoryListGetRequest() *AlibabaAscpChannelDistributorInventoryListGetAPIRequest {
	return &AlibabaAscpChannelDistributorInventoryListGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAscpChannelDistributorInventoryListGetAPIRequest) Reset() {
	r._inventoryRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpChannelDistributorInventoryListGetAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.channel.distributor.inventory.list.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpChannelDistributorInventoryListGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscpChannelDistributorInventoryListGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInventoryRequest is InventoryRequest Setter
// 系统自动生成
func (r *AlibabaAscpChannelDistributorInventoryListGetAPIRequest) SetInventoryRequest(_inventoryRequest *BatchChannelInventoryQuery) error {
	r._inventoryRequest = _inventoryRequest
	r.Set("inventory_request", _inventoryRequest)
	return nil
}

// GetInventoryRequest InventoryRequest Getter
func (r AlibabaAscpChannelDistributorInventoryListGetAPIRequest) GetInventoryRequest() *BatchChannelInventoryQuery {
	return r._inventoryRequest
}

var poolAlibabaAscpChannelDistributorInventoryListGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAscpChannelDistributorInventoryListGetRequest()
	},
}

// GetAlibabaAscpChannelDistributorInventoryListGetRequest 从 sync.Pool 获取 AlibabaAscpChannelDistributorInventoryListGetAPIRequest
func GetAlibabaAscpChannelDistributorInventoryListGetAPIRequest() *AlibabaAscpChannelDistributorInventoryListGetAPIRequest {
	return poolAlibabaAscpChannelDistributorInventoryListGetAPIRequest.Get().(*AlibabaAscpChannelDistributorInventoryListGetAPIRequest)
}

// ReleaseAlibabaAscpChannelDistributorInventoryListGetAPIRequest 将 AlibabaAscpChannelDistributorInventoryListGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaAscpChannelDistributorInventoryListGetAPIRequest(v *AlibabaAscpChannelDistributorInventoryListGetAPIRequest) {
	v.Reset()
	poolAlibabaAscpChannelDistributorInventoryListGetAPIRequest.Put(v)
}
