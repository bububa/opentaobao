package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomeprojectbuildingecodeupdateAPIRequest 新房楼栋修改e码 API请求
// alibaba.alihouse.newhome.project.building.ecode.update
//
// 新房楼栋修改e码
type AlibabaalihousenewhomeprojectbuildingecodeupdateAPIRequest struct {
	model.Params
	// 楼栋请求实体
	_updateBuilding *UpdateEcodeBuildingDto
}

// NewAlibabaalihousenewhomeprojectbuildingecodeupdateRequest 初始化AlibabaalihousenewhomeprojectbuildingecodeupdateAPIRequest对象
func NewAlibabaalihousenewhomeprojectbuildingecodeupdateRequest() *AlibabaalihousenewhomeprojectbuildingecodeupdateAPIRequest {
	return &AlibabaalihousenewhomeprojectbuildingecodeupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousenewhomeprojectbuildingecodeupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.project.building.ecode.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousenewhomeprojectbuildingecodeupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousenewhomeprojectbuildingecodeupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUpdateBuilding is UpdateBuilding Setter
// 楼栋请求实体
func (r *AlibabaalihousenewhomeprojectbuildingecodeupdateAPIRequest) SetUpdateBuilding(_updateBuilding *UpdateEcodeBuildingDto) error {
	r._updateBuilding = _updateBuilding
	r.Set("update_building", _updateBuilding)
	return nil
}

// GetUpdateBuilding UpdateBuilding Getter
func (r AlibabaalihousenewhomeprojectbuildingecodeupdateAPIRequest) GetUpdateBuilding() *UpdateEcodeBuildingDto {
	return r._updateBuilding
}
