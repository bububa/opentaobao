package inventory

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLocationRelationQueryAPIRequest 地点关联关系查询 API请求
// taobao.location.relation.query
//
// 地点关联关系查询
// 门店和仓库关联关系查询
type TaobaoLocationRelationQueryAPIRequest struct {
	model.Params
	// 关系查询
	_locationRelation *LocationRelationDto
}

// NewTaobaoLocationRelationQueryRequest 初始化TaobaoLocationRelationQueryAPIRequest对象
func NewTaobaoLocationRelationQueryRequest() *TaobaoLocationRelationQueryAPIRequest {
	return &TaobaoLocationRelationQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoLocationRelationQueryAPIRequest) Reset() {
	r._locationRelation = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLocationRelationQueryAPIRequest) GetApiMethodName() string {
	return "taobao.location.relation.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLocationRelationQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoLocationRelationQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLocationRelation is LocationRelation Setter
// 关系查询
func (r *TaobaoLocationRelationQueryAPIRequest) SetLocationRelation(_locationRelation *LocationRelationDto) error {
	r._locationRelation = _locationRelation
	r.Set("location_relation", _locationRelation)
	return nil
}

// GetLocationRelation LocationRelation Getter
func (r TaobaoLocationRelationQueryAPIRequest) GetLocationRelation() *LocationRelationDto {
	return r._locationRelation
}

var poolTaobaoLocationRelationQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoLocationRelationQueryRequest()
	},
}

// GetTaobaoLocationRelationQueryRequest 从 sync.Pool 获取 TaobaoLocationRelationQueryAPIRequest
func GetTaobaoLocationRelationQueryAPIRequest() *TaobaoLocationRelationQueryAPIRequest {
	return poolTaobaoLocationRelationQueryAPIRequest.Get().(*TaobaoLocationRelationQueryAPIRequest)
}

// ReleaseTaobaoLocationRelationQueryAPIRequest 将 TaobaoLocationRelationQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoLocationRelationQueryAPIRequest(v *TaobaoLocationRelationQueryAPIRequest) {
	v.Reset()
	poolTaobaoLocationRelationQueryAPIRequest.Put(v)
}
