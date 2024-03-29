package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeProjectBuildingEcodeUpdateAPIRequest 新房楼栋修改e码 API请求
// alibaba.alihouse.newhome.project.building.ecode.update
//
// 新房楼栋修改e码
type AlibabaAlihouseNewhomeProjectBuildingEcodeUpdateAPIRequest struct {
	model.Params
	// 楼栋请求实体
	_updateBuilding *UpdateECodeBuildingDto
}

// NewAlibabaAlihouseNewhomeProjectBuildingEcodeUpdateRequest 初始化AlibabaAlihouseNewhomeProjectBuildingEcodeUpdateAPIRequest对象
func NewAlibabaAlihouseNewhomeProjectBuildingEcodeUpdateRequest() *AlibabaAlihouseNewhomeProjectBuildingEcodeUpdateAPIRequest {
	return &AlibabaAlihouseNewhomeProjectBuildingEcodeUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseNewhomeProjectBuildingEcodeUpdateAPIRequest) Reset() {
	r._updateBuilding = nil
	r.Params.ToZero()
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
func (r *AlibabaAlihouseNewhomeProjectBuildingEcodeUpdateAPIRequest) SetUpdateBuilding(_updateBuilding *UpdateECodeBuildingDto) error {
	r._updateBuilding = _updateBuilding
	r.Set("update_building", _updateBuilding)
	return nil
}

// GetUpdateBuilding UpdateBuilding Getter
func (r AlibabaAlihouseNewhomeProjectBuildingEcodeUpdateAPIRequest) GetUpdateBuilding() *UpdateECodeBuildingDto {
	return r._updateBuilding
}

var poolAlibabaAlihouseNewhomeProjectBuildingEcodeUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseNewhomeProjectBuildingEcodeUpdateRequest()
	},
}

// GetAlibabaAlihouseNewhomeProjectBuildingEcodeUpdateRequest 从 sync.Pool 获取 AlibabaAlihouseNewhomeProjectBuildingEcodeUpdateAPIRequest
func GetAlibabaAlihouseNewhomeProjectBuildingEcodeUpdateAPIRequest() *AlibabaAlihouseNewhomeProjectBuildingEcodeUpdateAPIRequest {
	return poolAlibabaAlihouseNewhomeProjectBuildingEcodeUpdateAPIRequest.Get().(*AlibabaAlihouseNewhomeProjectBuildingEcodeUpdateAPIRequest)
}

// ReleaseAlibabaAlihouseNewhomeProjectBuildingEcodeUpdateAPIRequest 将 AlibabaAlihouseNewhomeProjectBuildingEcodeUpdateAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseNewhomeProjectBuildingEcodeUpdateAPIRequest(v *AlibabaAlihouseNewhomeProjectBuildingEcodeUpdateAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseNewhomeProjectBuildingEcodeUpdateAPIRequest.Put(v)
}
