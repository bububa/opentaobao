package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpchanneldistributorinventorylistgetAPIRequest 批量查询渠道库存 API请求
// alibaba.ascp.channel.distributor.inventory.list.get
//
// 淘外分销批量查询渠道产品的库存
type AlibabaascpchanneldistributorinventorylistgetAPIRequest struct {
	model.Params
	// 系统自动生成
	_inventoryRequest *BatchChannelInventoryQuery
}

// NewAlibabaascpchanneldistributorinventorylistgetRequest 初始化AlibabaascpchanneldistributorinventorylistgetAPIRequest对象
func NewAlibabaascpchanneldistributorinventorylistgetRequest() *AlibabaascpchanneldistributorinventorylistgetAPIRequest {
	return &AlibabaascpchanneldistributorinventorylistgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaascpchanneldistributorinventorylistgetAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.channel.distributor.inventory.list.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaascpchanneldistributorinventorylistgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaascpchanneldistributorinventorylistgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInventoryRequest is InventoryRequest Setter
// 系统自动生成
func (r *AlibabaascpchanneldistributorinventorylistgetAPIRequest) SetInventoryRequest(_inventoryRequest *BatchChannelInventoryQuery) error {
	r._inventoryRequest = _inventoryRequest
	r.Set("inventory_request", _inventoryRequest)
	return nil
}

// GetInventoryRequest InventoryRequest Getter
func (r AlibabaascpchanneldistributorinventorylistgetAPIRequest) GetInventoryRequest() *BatchChannelInventoryQuery {
	return r._inventoryRequest
}
