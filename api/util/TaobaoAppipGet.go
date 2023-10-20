package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// TaobaoAppipGet 获取ISV发起请求服务器IP
// taobao.appip.get
//
// 获取ISV发起请求服务器IP
func TaobaoAppipGet(clt *core.SDKClient, req *util.TaobaoAppipGetAPIRequest, resp *util.TaobaoAppipGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
