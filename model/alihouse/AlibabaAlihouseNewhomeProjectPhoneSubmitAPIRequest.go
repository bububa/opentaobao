package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihouseNewhomeProjectPhoneSubmitAPIRequest
提交楼盘电话 API请求
alibaba.alihouse.newhome.project.phone.submit

提交楼盘电话 */
type AlibabaAlihouseNewhomeProjectPhoneSubmitAPIRequest struct {
	model.Params
	// 楼盘电话
	_projectPhoneDto *ProjectPhoneDto
}

// New
