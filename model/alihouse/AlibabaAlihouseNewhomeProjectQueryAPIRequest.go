package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihouseNewhomeProjectQueryAPIRequest
查询楼盘相关信息 API请求
alibaba.alihouse.newhome.project.query

根据outerid查询楼盘相关信息 */
type AlibabaAlihouseNewhomeProjectQueryAPIRequest struct {
	model.Params
	// 外部楼盘/小区id
	_outerId string
}

// New
