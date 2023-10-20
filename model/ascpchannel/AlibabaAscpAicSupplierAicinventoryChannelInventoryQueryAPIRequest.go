package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpAicSupplierAicinventoryChannelInventoryQueryAPIRequest 供应商直发-商家仓库存查询服务 API请求
// alibaba.ascp.aic.supplier.aicinventory.channel.inventory.query
//
// 提供商家基于货品、供应商、仓，查询ascp 实时商家仓库存查询数据。
type AlibabaAscpAicSupplierAicinventoryChannelInventoryQueryAPIRequest struct {
	model.Params
	// 商家仓库存查询请求参数
	_merchantInventoryQueryRequest *MerchantInventoryQuery
}

// NewAlibabaAscpAicSupplierAicinventoryChannelInventoryQueryRequest 初始化AlibabaAscpAicSupplierAicinventoryChannelInventoryQueryAPIRequest对象
func NewAlibabaAscpAicSupplierAicinventoryChannelInventoryQueryRequest() *AlibabaAscpAicSupplierAicinventoryChannelInventoryQueryAPIRequest {
	return &AlibabaAscpAicSupplierAicinventoryChannelInventoryQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpAicSupplierAicinventoryChannelInventoryQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.aic.supplier.aicinventory.channel.inventory.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpAicSupplierAicinventoryChannelInventoryQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscpAicSupplierAicinventoryChannelInventoryQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMerchantInventoryQueryRequest is MerchantInventoryQueryRequest Setter
// 商家仓库存查询请求参数
func (r *AlibabaAscpAicSupplierAicinventoryChannelInventoryQueryAPIRequest) SetMerchantInventoryQueryRequest(_merchantInventoryQueryRequest *MerchantInventoryQuery) error {
	r._merchantInventoryQueryRequest = _merchantInventoryQueryRequest
	r.Set("merchant_inventory_query_request", _merchantInventoryQueryRequest)
	return nil
}

// GetMerchantInventoryQueryRequest MerchantInventoryQueryRequest Getter
func (r AlibabaAscpAicSupplierAicinventoryChannelInventoryQueryAPIRequest) GetMerchantInventoryQueryRequest() *MerchantInventoryQuery {
	return r._merchantInventoryQueryRequest
}
