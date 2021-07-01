package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihouseNewhomeProjectSubmitAPIRequest
提交楼盘信息 API请求
alibaba.alihouse.newhome.project.submit

提交楼盘信息 */
type AlibabaAlihouseNewhomeProjectSubmitAPIRequest struct {
	model.Params
	// 楼盘对象
	_ebbasProjectDto *EbbasProjectDto
}

// New
