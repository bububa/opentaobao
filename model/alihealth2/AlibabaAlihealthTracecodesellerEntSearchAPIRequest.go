package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthTracecodesellerEntSearchAPIRequest
查询商家信息 API请求
alibaba.alihealth.tracecodeseller.ent.search

查询商家信息 */
type AlibabaAlihealthTracecodesellerEntSearchAPIRequest struct {
	model.Params
	// appkey
	_skeyCode string
	// 商家名称
	_name string
	// 淘宝名
	_tbUserId string
}

// New
