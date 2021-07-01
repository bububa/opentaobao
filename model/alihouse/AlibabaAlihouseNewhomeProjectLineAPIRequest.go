package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihouseNewhomeProjectLineAPIRequest
楼盘上下架 API请求
alibaba.alihouse.newhome.project.line

上下架楼盘 */
type AlibabaAlihouseNewhomeProjectLineAPIRequest struct {
	model.Params
	// 外部id
	_outerId string
	// 0-下架 1-上架
	_type *model.File
}

// New
