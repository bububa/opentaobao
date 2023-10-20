package ascpffo

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressascponwayinventoryqueryAPIRequest AliExpress在途库存查询API API请求
// aliexpress.ascp.onway.inventory.query
//
// AliExpress在途库存查询API
type AliexpressascponwayinventoryqueryAPIRequest struct {
	model.Params
	// 查询DTO
	_onWayInventoryQueryDto *OnWayInventoryQueryDto
}

// NewAliexpressascponwayinventoryqueryRequest 初始化AliexpressascponwayinventoryqueryAPIRequest对象
func NewAliexpressascponwayinventoryqueryRequest() *AliexpressascponwayinventoryqueryAPIRequest {
	return &AliexpressascponwayinventoryqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressascponwayinventoryqueryAPIRequest) GetApiMethodName() string {
	return "aliexpress.ascp.onway.inventory.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressascponwayinventoryqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressascponwayinventoryqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOnWayInventoryQueryDto is OnWayInventoryQueryDto Setter
// 查询DTO
func (r *AliexpressascponwayinventoryqueryAPIRequest) SetOnWayInventoryQueryDto(_onWayInventoryQueryDto *OnWayInventoryQueryDto) error {
	r._onWayInventoryQueryDto = _onWayInventoryQueryDto
	r.Set("on_way_inventory_query_dto", _onWayInventoryQueryDto)
	return nil
}

// GetOnWayInventoryQueryDto OnWayInventoryQueryDto Getter
func (r AliexpressascponwayinventoryqueryAPIRequest) GetOnWayInventoryQueryDto() *OnWayInventoryQueryDto {
	return r._onWayInventoryQueryDto
}
