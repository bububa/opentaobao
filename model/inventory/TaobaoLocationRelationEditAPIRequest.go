package inventory

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaolocationrelationeditAPIRequest 地点关联关系增量编辑 API请求
// taobao.location.relation.edit
//
// 地点关联关系增量编辑
type TaobaolocationrelationeditAPIRequest struct {
	model.Params
	// 关系对象列表
	_locationRelationList []LocationRelationDto
}

// NewTaobaolocationrelationeditRequest 初始化TaobaolocationrelationeditAPIRequest对象
func NewTaobaolocationrelationeditRequest() *TaobaolocationrelationeditAPIRequest {
	return &TaobaolocationrelationeditAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaolocationrelationeditAPIRequest) GetApiMethodName() string {
	return "taobao.location.relation.edit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaolocationrelationeditAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaolocationrelationeditAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLocationRelationList is LocationRelationList Setter
// 关系对象列表
func (r *TaobaolocationrelationeditAPIRequest) SetLocationRelationList(_locationRelationList []LocationRelationDto) error {
	r._locationRelationList = _locationRelationList
	r.Set("location_relation_list", _locationRelationList)
	return nil
}

// GetLocationRelationList LocationRelationList Getter
func (r TaobaolocationrelationeditAPIRequest) GetLocationRelationList() []LocationRelationDto {
	return r._locationRelationList
}
