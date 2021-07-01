package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihouseNewhomeProjectAdviserSubmitAPIRequest
提交楼盘顾问 API请求
alibaba.alihouse.newhome.project.adviser.submit

提交楼盘顾问 */
type AlibabaAlihouseNewhomeProjectAdviserSubmitAPIRequest struct {
	model.Params
	// 顾问列表
	_advisers []ProjectAdviserDto
}

// New
