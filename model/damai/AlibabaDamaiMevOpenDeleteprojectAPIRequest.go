package damai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMevOpenDeleteprojectAPIRequest 大麦换验平台-第三方对外开放-项目接口deleteProject API请求
// alibaba.damai.mev.open.deleteproject
//
// deleteProject
type AlibabaDamaiMevOpenDeleteprojectAPIRequest struct {
	model.Params
	// 入参deleteProjectParam
	_deleteProjectParam *ProjectIdOpenParam
}

// NewAlibabaDamaiMevOpenDeleteprojectRequest 初始化AlibabaDamaiMevOpenDeleteprojectAPIRequest对象
func NewAlibabaDamaiMevOpenDeleteprojectRequest() *AlibabaDamaiMevOpenDeleteprojectAPIRequest {
	return &AlibabaDamaiMevOpenDeleteprojectAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMevOpenDeleteprojectAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.mev.open.deleteproject"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMevOpenDeleteprojectAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDamaiMevOpenDeleteprojectAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeleteProjectParam is DeleteProjectParam Setter
// 入参deleteProjectParam
func (r *AlibabaDamaiMevOpenDeleteprojectAPIRequest) SetDeleteProjectParam(_deleteProjectParam *ProjectIdOpenParam) error {
	r._deleteProjectParam = _deleteProjectParam
	r.Set("delete_project_param", _deleteProjectParam)
	return nil
}

// GetDeleteProjectParam DeleteProjectParam Getter
func (r AlibabaDamaiMevOpenDeleteprojectAPIRequest) GetDeleteProjectParam() *ProjectIdOpenParam {
	return r._deleteProjectParam
}
