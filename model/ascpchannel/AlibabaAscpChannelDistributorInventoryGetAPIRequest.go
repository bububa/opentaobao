package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpChannelDistributorInventoryGetAPIRequest 链渠道中心淘外库存查询 API请求
// alibaba.ascp.channel.distributor.inventory.get
//
// 此api为淘外分销的渠道产品库存查询标准api，淘外分销商专用
type AlibabaAscpChannelDistributorInventoryGetAPIRequest struct {
	model.Params
	// 入参
	_inventoryRequest *ChannelInventoryQuery
}

// NewAlibabaAscpChannelDistributorInventoryGetRequest 初始化AlibabaAscpChannelDistributorInventoryGetAPIRequest对象
func NewAlibabaAscpChannelDistributorInventoryGetRequest() *AlibabaAscpChannelDistributorInventoryGetAPIRequest {
	return &AlibabaAscpChannelDistributorInventoryGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpChannelDistributorInventoryGetAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.channel.distributor.inventory.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpChannelDistributorInventoryGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscpChannelDistributorInventoryGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInventoryRequest is InventoryRequest Setter
// 入参
func (r *AlibabaAscpChannelDistributorInventoryGetAPIRequest) SetInventoryRequest(_inventoryRequest *ChannelInventoryQuery) error {
	r._inventoryRequest = _inventoryRequest
	r.Set("inventory_request", _inventoryRequest)
	return nil
}

// GetInventoryRequest InventoryRequest Getter
func (r AlibabaAscpChannelDistributorInventoryGetAPIRequest) GetInventoryRequest() *ChannelInventoryQuery {
	return r._inventoryRequest
}
