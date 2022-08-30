package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseVrfactoryProductionSyncAPIRequest vr生产数据上翻 API请求
// alibaba.alihouse.vrfactory.production.sync
//
// vr生产数据上翻
type AlibabaAlihouseVrfactoryProductionSyncAPIRequest struct {
	model.Params
	// 参数
	_projectVrBuildDataDto *ProjectVrBuildDataDto
}

// NewAlibabaAlihouseVrfactoryProductionSyncRequest 初始化AlibabaAlihouseVrfactoryProductionSyncAPIRequest对象
func NewAlibabaAlihouseVrfactoryProductionSyncRequest() *AlibabaAlihouseVrfactoryProductionSyncAPIRequest {
	return &AlibabaAlihouseVrfactoryProductionSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseVrfactoryProductionSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.vrfactory.production.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseVrfactoryProductionSyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetProjectVrBuildDataDto is ProjectVrBuildDataDto Setter
// 参数
func (r *AlibabaAlihouseVrfactoryProductionSyncAPIRequest) SetProjectVrBuildDataDto(_projectVrBuildDataDto *ProjectVrBuildDataDto) error {
	r._projectVrBuildDataDto = _projectVrBuildDataDto
	r.Set("project_vr_build_data_dto", _projectVrBuildDataDto)
	return nil
}

// GetProjectVrBuildDataDto ProjectVrBuildDataDto Getter
func (r AlibabaAlihouseVrfactoryProductionSyncAPIRequest) GetProjectVrBuildDataDto() *ProjectVrBuildDataDto {
	return r._projectVrBuildDataDto
}
