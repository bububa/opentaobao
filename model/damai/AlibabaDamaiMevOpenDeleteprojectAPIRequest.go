package damai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDamaiMevOpenDeleteprojectAPIRequest
大麦换验平台-第三方对外开放-项目接口deleteProject API请求
alibaba.damai.mev.open.deleteproject

deleteProject */
type AlibabaDamaiMevOpenDeleteprojectAPIRequest struct {
	model.Params
	// 入参deleteProjectParam
	_deleteProjectParam *ProjectIdOpenParam
}

// New
