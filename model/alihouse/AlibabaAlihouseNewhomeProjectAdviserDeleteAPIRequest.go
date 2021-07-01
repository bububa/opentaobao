package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihouseNewhomeProjectAdviserDeleteAPIRequest
删除楼盘顾问 API请求
alibaba.alihouse.newhome.project.adviser.delete

删除楼盘顾问 */
type AlibabaAlihouseNewhomeProjectAdviserDeleteAPIRequest struct {
	model.Params
	// 外部顾问ID
	_outerConsultantId string
}

// New
