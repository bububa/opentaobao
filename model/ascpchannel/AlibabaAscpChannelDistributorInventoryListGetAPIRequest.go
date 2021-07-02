package ascpchannel

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpChannelDistributorInventoryListGetAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.channel.distributor.inventory.list.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpChannelDistributorInventoryListGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
