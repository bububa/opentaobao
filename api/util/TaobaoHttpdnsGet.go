package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// TaobaoHttpdnsGet TOPDNS配置
// taobao.httpdns.get
//
// 获取TOP DNS配置
func TaobaoHttpdnsGet(clt *core.SDKClient, req *util.TaobaoHttpdnsGetAPIRequest, resp *util.TaobaoHttpdnsGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
