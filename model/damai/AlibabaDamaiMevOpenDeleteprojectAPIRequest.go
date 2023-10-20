package damai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadamaimevopendeleteprojectAPIRequest 大麦换验平台-第三方对外开放-项目接口deleteProject API请求
// alibaba.damai.mev.open.deleteproject
//
// deleteProject
type AlibabadamaimevopendeleteprojectAPIRequest struct {
	model.Params
	// 入参deleteProjectParam
	_deleteProjectParam *ProjectIdOpenParam
}

// NewAlibabadamaimevopendeleteprojectRequest 初始化AlibabadamaimevopendeleteprojectAPIRequest对象
func NewAlibabadamaimevopendeleteprojectRequest() *AlibabadamaimevopendeleteprojectAPIRequest {
	return &AlibabadamaimevopendeleteprojectAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadamaimevopendeleteprojectAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.mev.open.deleteproject"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadamaimevopendeleteprojectAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadamaimevopendeleteprojectAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeleteProjectParam is DeleteProjectParam Setter
// 入参deleteProjectParam
func (r *AlibabadamaimevopendeleteprojectAPIRequest) SetDeleteProjectParam(_deleteProjectParam *ProjectIdOpenParam) error {
	r._deleteProjectParam = _deleteProjectParam
	r.Set("delete_project_param", _deleteProjectParam)
	return nil
}

// GetDeleteProjectParam DeleteProjectParam Getter
func (r AlibabadamaimevopendeleteprojectAPIRequest) GetDeleteProjectParam() *ProjectIdOpenParam {
	return r._deleteProjectParam
}
