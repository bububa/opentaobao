package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihouseNewhomeProjectDynamicDeleteAPIRequest
楼盘快讯删除 API请求
alibaba.alihouse.newhome.project.dynamic.delete

楼盘快讯删除 */
type AlibabaAlihouseNewhomeProjectDynamicDeleteAPIRequest struct {
	model.Params
	// 外部动态ID
	_outerDynamicId string
}

// New
