package inventory

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaolocationrelationqueryAPIRequest 地点关联关系查询 API请求
// taobao.location.relation.query
//
// 地点关联关系查询
// 门店和仓库关联关系查询
type TaobaolocationrelationqueryAPIRequest struct {
	model.Params
	// 关系查询
	_locationRelation *LocationRelationDto
}

// NewTaobaolocationrelationqueryRequest 初始化TaobaolocationrelationqueryAPIRequest对象
func NewTaobaolocationrelationqueryRequest() *TaobaolocationrelationqueryAPIRequest {
	return &TaobaolocationrelationqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaolocationrelationqueryAPIRequest) GetApiMethodName() string {
	return "taobao.location.relation.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaolocationrelationqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaolocationrelationqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLocationRelation is LocationRelation Setter
// 关系查询
func (r *TaobaolocationrelationqueryAPIRequest) SetLocationRelation(_locationRelation *LocationRelationDto) error {
	r._locationRelation = _locationRelation
	r.Set("location_relation", _locationRelation)
	return nil
}

// GetLocationRelation LocationRelation Getter
func (r TaobaolocationrelationqueryAPIRequest) GetLocationRelation() *LocationRelationDto {
	return r._locationRelation
}
