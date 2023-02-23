package moscm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosGoodsInventoryGetinventorysAPIRequest 可售库存查询 API请求
// alibaba.mos.goods.inventory.getinventorys
//
// 查询商品的可售、在库和占库数量
type AlibabaMosGoodsInventoryGetinventorysAPIRequest struct {
	model.Params
	// 查询对象
	_paramVirtualInventoryQueryDto *VirtualInventoryQueryDto
}

// NewAlibabaMosGoodsInventoryGetinventorysRequest 初始化AlibabaMosGoodsInventoryGetinventorysAPIRequest对象
func NewAlibabaMosGoodsInventoryGetinventorysRequest() *AlibabaMosGoodsInventoryGetinventorysAPIRequest {
	return &AlibabaMosGoodsInventoryGetinventorysAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMosGoodsInventoryGetinventorysAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.goods.inventory.getinventorys"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMosGoodsInventoryGetinventorysAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMosGoodsInventoryGetinventorysAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamVirtualInventoryQueryDto is ParamVirtualInventoryQueryDto Setter
// 查询对象
func (r *AlibabaMosGoodsInventoryGetinventorysAPIRequest) SetParamVirtualInventoryQueryDto(_paramVirtualInventoryQueryDto *VirtualInventoryQueryDto) error {
	r._paramVirtualInventoryQueryDto = _paramVirtualInventoryQueryDto
	r.Set("param_virtual_inventory_query_dto", _paramVirtualInventoryQueryDto)
	return nil
}

// GetParamVirtualInventoryQueryDto ParamVirtualInventoryQueryDto Getter
func (r AlibabaMosGoodsInventoryGetinventorysAPIRequest) GetParamVirtualInventoryQueryDto() *VirtualInventoryQueryDto {
	return r._paramVirtualInventoryQueryDto
}
