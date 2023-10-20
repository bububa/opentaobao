package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeProjectUpdateItemInfoAPIRequest 更新楼盘商品信息 API请求
// alibaba.alihouse.newhome.project.update.item.info
//
// 更新楼盘商品信息
type AlibabaAlihouseNewhomeProjectUpdateItemInfoAPIRequest struct {
	model.Params
	// 实体类
	_projectUpdateItemInfoDto *ProjectUpdateItemInfoDto
}

// NewAlibabaAlihouseNewhomeProjectUpdateItemInfoRequest 初始化AlibabaAlihouseNewhomeProjectUpdateItemInfoAPIRequest对象
func NewAlibabaAlihouseNewhomeProjectUpdateItemInfoRequest() *AlibabaAlihouseNewhomeProjectUpdateItemInfoAPIRequest {
	return &AlibabaAlihouseNewhomeProjectUpdateItemInfoAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseNewhomeProjectUpdateItemInfoAPIRequest) Reset() {
	r._projectUpdateItemInfoDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeProjectUpdateItemInfoAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.project.update.item.info"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeProjectUpdateItemInfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeProjectUpdateItemInfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProjectUpdateItemInfoDto is ProjectUpdateItemInfoDto Setter
// 实体类
func (r *AlibabaAlihouseNewhomeProjectUpdateItemInfoAPIRequest) SetProjectUpdateItemInfoDto(_projectUpdateItemInfoDto *ProjectUpdateItemInfoDto) error {
	r._projectUpdateItemInfoDto = _projectUpdateItemInfoDto
	r.Set("project_update_item_info_dto", _projectUpdateItemInfoDto)
	return nil
}

// GetProjectUpdateItemInfoDto ProjectUpdateItemInfoDto Getter
func (r AlibabaAlihouseNewhomeProjectUpdateItemInfoAPIRequest) GetProjectUpdateItemInfoDto() *ProjectUpdateItemInfoDto {
	return r._projectUpdateItemInfoDto
}

var poolAlibabaAlihouseNewhomeProjectUpdateItemInfoAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseNewhomeProjectUpdateItemInfoRequest()
	},
}

// GetAlibabaAlihouseNewhomeProjectUpdateItemInfoRequest 从 sync.Pool 获取 AlibabaAlihouseNewhomeProjectUpdateItemInfoAPIRequest
func GetAlibabaAlihouseNewhomeProjectUpdateItemInfoAPIRequest() *AlibabaAlihouseNewhomeProjectUpdateItemInfoAPIRequest {
	return poolAlibabaAlihouseNewhomeProjectUpdateItemInfoAPIRequest.Get().(*AlibabaAlihouseNewhomeProjectUpdateItemInfoAPIRequest)
}

// ReleaseAlibabaAlihouseNewhomeProjectUpdateItemInfoAPIRequest 将 AlibabaAlihouseNewhomeProjectUpdateItemInfoAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseNewhomeProjectUpdateItemInfoAPIRequest(v *AlibabaAlihouseNewhomeProjectUpdateItemInfoAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseNewhomeProjectUpdateItemInfoAPIRequest.Put(v)
}
