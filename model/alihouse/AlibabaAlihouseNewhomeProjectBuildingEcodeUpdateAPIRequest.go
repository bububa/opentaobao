package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeProjectBuildingEcodeUpdateAPIRequest 新房楼栋修改e码 API请求
// alibaba.alihouse.newhome.project.building.ecode.update
//
// 新房楼栋修改e码
type AlibabaAlihouseNewhomeProjectBuildingEcodeUpdateAPIRequest struct {
	model.Params
	// 楼栋请求实体
	_updateBuilding *UpdateEcodeBuildingDto
}

// NewAlibabaAlihouseNewhomeProjectBuildingEcodeUpdateRequest 初始化AlibabaAlihouseNewhomeProjectBuildingEcodeUpdateAPIRequest对象
func NewAlibabaAlihouseNewhomeProjectBuildingEcodeUpdateRequest() *AlibabaAlihouseNewhomeProjectBuildingEcodeUpdateAPIRequest {
	return &AlibabaAlihouseNewhomeProjectBuildingEcodeUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeProjectBuildingEcodeUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.project.building.ecode.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeProjectBuildingEcodeUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeProjectBuildingEcodeUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUpdateBuilding is UpdateBuilding Setter
// 楼栋请求实体
func (r *AlibabaAlihouseNewhomeProjectBuildingEcodeUpdateAPIRequest) SetUpdateBuilding(_updateBuilding *UpdateEcodeBuildingDto) error {
	r._updateBuilding = _updateBuilding
	r.Set("update_building", _updateBuilding)
	return nil
}

// GetUpdateBuilding UpdateBuilding Getter
func (r AlibabaAlihouseNewhomeProjectBuildingEcodeUpdateAPIRequest) GetUpdateBuilding() *UpdateEcodeBuildingDto {
	return r._updateBuilding
}
