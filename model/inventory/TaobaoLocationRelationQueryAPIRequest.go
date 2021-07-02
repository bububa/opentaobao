package inventory

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLocationRelationQueryAPIRequest) GetApiMethodName() string {
	return "taobao.location.relation.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLocationRelationQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is LocationRelation Setter
// 关系查询
func (r *TaobaoLocationRelationQueryAPIRequest) SetLocationRelation(_locationRelation *LocationRelationDto) error {
	r._locationRelation = _locationRelation
	r.Set("location_relation", _locationRelation)
	return nil
}

// Get LocationRelation Getter
func (r TaobaoLocationRelationQueryAPIRequest) GetLocationRelation() *LocationRelationDto {
	return r._locationRelation
}
