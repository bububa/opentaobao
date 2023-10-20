package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomeprojectbuildingAPIRequest 楼栋同步 API请求
// alibaba.alihouse.newhome.project.building
//
// 楼栋同步
type AlibabaalihousenewhomeprojectbuildingAPIRequest struct {
	model.Params
	// 楼栋请求实体
	_buildingDto *BuildingDto
}

// NewAlibabaalihousenewhomeprojectbuildingRequest 初始化AlibabaalihousenewhomeprojectbuildingAPIRequest对象
func NewAlibabaalihousenewhomeprojectbuildingRequest() *AlibabaalihousenewhomeprojectbuildingAPIRequest {
	return &AlibabaalihousenewhomeprojectbuildingAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousenewhomeprojectbuildingAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.project.building"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousenewhomeprojectbuildingAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousenewhomeprojectbuildingAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBuildingDto is BuildingDto Setter
// 楼栋请求实体
func (r *AlibabaalihousenewhomeprojectbuildingAPIRequest) SetBuildingDto(_buildingDto *BuildingDto) error {
	r._buildingDto = _buildingDto
	r.Set("building_dto", _buildingDto)
	return nil
}

// GetBuildingDto BuildingDto Getter
func (r AlibabaalihousenewhomeprojectbuildingAPIRequest) GetBuildingDto() *BuildingDto {
	return r._buildingDto
}
