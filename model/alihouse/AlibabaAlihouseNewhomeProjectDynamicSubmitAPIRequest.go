package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihouseNewhomeProjectDynamicSubmitAPIRequest
提交楼盘快讯 API请求
alibaba.alihouse.newhome.project.dynamic.submit

提交楼盘快讯 */
type AlibabaAlihouseNewhomeProjectDynamicSubmitAPIRequest struct {
	model.Params
	// 楼盘动态列表
	_projectDynamics []ProjectDynamicDto
}

// New
