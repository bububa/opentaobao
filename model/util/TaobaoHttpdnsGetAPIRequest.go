package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoHttpdnsGetAPIRequest
TOPDNS配置 API请求
taobao.httpdns.get

获取TOP DNS配置 */
type TaobaoHttpdnsGetAPIRequest struct {
	model.Params
}

// New
