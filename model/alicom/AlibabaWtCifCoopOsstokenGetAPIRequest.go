package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWtCifCoopOsstokenGetAPIRequest
获取oss签名接口 API请求
alibaba.wt.cif.coop.osstoken.get

商家合作上传oss图片获取token接口 */
type AlibabaWtCifCoopOsstokenGetAPIRequest struct {
	model.Params
	// 调用方的应用名
	_appName string
	// 系统分配的source
	_source string
	// 系统分配的biz
	_biz string
}

// New
