package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihouseNewhomeProjectCooperationSubmitAPIRequest
提交KA合作楼盘 API请求
alibaba.alihouse.newhome.project.cooperation.submit

提交KA合作楼盘 */
type AlibabaAlihouseNewhomeProjectCooperationSubmitAPIRequest struct {
	model.Params
	// ka合作对象
	_projectCooperationDto *ProjectCooperationDto
}

// New
