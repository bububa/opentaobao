package moscm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamosgoodsinventorygetinventorysAPIRequest 可售库存查询 API请求
// alibaba.mos.goods.inventory.getinventorys
//
// 查询商品的可售、在库和占库数量
type AlibabamosgoodsinventorygetinventorysAPIRequest struct {
	model.Params
	// 查询对象
	_paramVirtualInventoryQueryDto *VirtualInventoryQueryDto
}

// NewAlibabamosgoodsinventorygetinventorysRequest 初始化AlibabamosgoodsinventorygetinventorysAPIRequest对象
func NewAlibabamosgoodsinventorygetinventorysRequest() *AlibabamosgoodsinventorygetinventorysAPIRequest {
	return &AlibabamosgoodsinventorygetinventorysAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamosgoodsinventorygetinventorysAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.goods.inventory.getinventorys"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamosgoodsinventorygetinventorysAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamosgoodsinventorygetinventorysAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamVirtualInventoryQueryDto is ParamVirtualInventoryQueryDto Setter
// 查询对象
func (r *AlibabamosgoodsinventorygetinventorysAPIRequest) SetParamVirtualInventoryQueryDto(_paramVirtualInventoryQueryDto *VirtualInventoryQueryDto) error {
	r._paramVirtualInventoryQueryDto = _paramVirtualInventoryQueryDto
	r.Set("param_virtual_inventory_query_dto", _paramVirtualInventoryQueryDto)
	return nil
}

// GetParamVirtualInventoryQueryDto ParamVirtualInventoryQueryDto Getter
func (r AlibabamosgoodsinventorygetinventorysAPIRequest) GetParamVirtualInventoryQueryDto() *VirtualInventoryQueryDto {
	return r._paramVirtualInventoryQueryDto
}
