package ascpffo

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressAscpOnwayInventoryQueryAPIRequest AliExpress在途库存查询API API请求
// aliexpress.ascp.onway.inventory.query
//
// AliExpress在途库存查询API
type AliexpressAscpOnwayInventoryQueryAPIRequest struct {
	model.Params
	// 查询DTO
	_onWayInventoryQueryDto *OnWayInventoryQueryDto
}

// NewAliexpressAscpOnwayInventoryQueryRequest 初始化AliexpressAscpOnwayInventoryQueryAPIRequest对象
func NewAliexpressAscpOnwayInventoryQueryRequest() *AliexpressAscpOnwayInventoryQueryAPIRequest {
	return &AliexpressAscpOnwayInventoryQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AliexpressAscpOnwayInventoryQueryAPIRequest) Reset() {
	r._onWayInventoryQueryDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressAscpOnwayInventoryQueryAPIRequest) GetApiMethodName() string {
	return "aliexpress.ascp.onway.inventory.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressAscpOnwayInventoryQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressAscpOnwayInventoryQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOnWayInventoryQueryDto is OnWayInventoryQueryDto Setter
// 查询DTO
func (r *AliexpressAscpOnwayInventoryQueryAPIRequest) SetOnWayInventoryQueryDto(_onWayInventoryQueryDto *OnWayInventoryQueryDto) error {
	r._onWayInventoryQueryDto = _onWayInventoryQueryDto
	r.Set("on_way_inventory_query_dto", _onWayInventoryQueryDto)
	return nil
}

// GetOnWayInventoryQueryDto OnWayInventoryQueryDto Getter
func (r AliexpressAscpOnwayInventoryQueryAPIRequest) GetOnWayInventoryQueryDto() *OnWayInventoryQueryDto {
	return r._onWayInventoryQueryDto
}

var poolAliexpressAscpOnwayInventoryQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAliexpressAscpOnwayInventoryQueryRequest()
	},
}

// GetAliexpressAscpOnwayInventoryQueryRequest 从 sync.Pool 获取 AliexpressAscpOnwayInventoryQueryAPIRequest
func GetAliexpressAscpOnwayInventoryQueryAPIRequest() *AliexpressAscpOnwayInventoryQueryAPIRequest {
	return poolAliexpressAscpOnwayInventoryQueryAPIRequest.Get().(*AliexpressAscpOnwayInventoryQueryAPIRequest)
}

// ReleaseAliexpressAscpOnwayInventoryQueryAPIRequest 将 AliexpressAscpOnwayInventoryQueryAPIRequest 放入 sync.Pool
func ReleaseAliexpressAscpOnwayInventoryQueryAPIRequest(v *AliexpressAscpOnwayInventoryQueryAPIRequest) {
	v.Reset()
	poolAliexpressAscpOnwayInventoryQueryAPIRequest.Put(v)
}
