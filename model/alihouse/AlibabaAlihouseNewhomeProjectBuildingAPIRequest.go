package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeProjectBuildingAPIRequest 楼栋同步 API请求
// alibaba.alihouse.newhome.project.building
//
// 楼栋同步
type AlibabaAlihouseNewhomeProjectBuildingAPIRequest struct {
	model.Params
	// 楼栋请求实体
	_buildingDto *BuildingDto
}

// NewAlibabaAlihouseNewhomeProjectBuildingRequest 初始化AlibabaAlihouseNewhomeProjectBuildingAPIRequest对象
func NewAlibabaAlihouseNewhomeProjectBuildingRequest() *AlibabaAlihouseNewhomeProjectBuildingAPIRequest {
	return &AlibabaAlihouseNewhomeProjectBuildingAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseNewhomeProjectBuildingAPIRequest) Reset() {
	r._buildingDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeProjectBuildingAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.project.building"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeProjectBuildingAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeProjectBuildingAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBuildingDto is BuildingDto Setter
// 楼栋请求实体
func (r *AlibabaAlihouseNewhomeProjectBuildingAPIRequest) SetBuildingDto(_buildingDto *BuildingDto) error {
	r._buildingDto = _buildingDto
	r.Set("building_dto", _buildingDto)
	return nil
}

// GetBuildingDto BuildingDto Getter
func (r AlibabaAlihouseNewhomeProjectBuildingAPIRequest) GetBuildingDto() *BuildingDto {
	return r._buildingDto
}

var poolAlibabaAlihouseNewhomeProjectBuildingAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseNewhomeProjectBuildingRequest()
	},
}

// GetAlibabaAlihouseNewhomeProjectBuildingRequest 从 sync.Pool 获取 AlibabaAlihouseNewhomeProjectBuildingAPIRequest
func GetAlibabaAlihouseNewhomeProjectBuildingAPIRequest() *AlibabaAlihouseNewhomeProjectBuildingAPIRequest {
	return poolAlibabaAlihouseNewhomeProjectBuildingAPIRequest.Get().(*AlibabaAlihouseNewhomeProjectBuildingAPIRequest)
}

// ReleaseAlibabaAlihouseNewhomeProjectBuildingAPIRequest 将 AlibabaAlihouseNewhomeProjectBuildingAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseNewhomeProjectBuildingAPIRequest(v *AlibabaAlihouseNewhomeProjectBuildingAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseNewhomeProjectBuildingAPIRequest.Put(v)
}
