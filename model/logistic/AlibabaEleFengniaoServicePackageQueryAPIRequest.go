package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEleFengniaoServicePackageQueryAPIRequest
预采购服务包查询接口 API请求
alibaba.ele.fengniao.service.package.query

查询门店所在经纬度可用服务包的接口 */
type AlibabaEleFengniaoServicePackageQueryAPIRequest struct {
	model.Params
	// 入参
	_param *Param
}

// New
