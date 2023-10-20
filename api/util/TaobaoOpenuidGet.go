package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// TaobaoOpenuidGet 获取授权账号对应的OpenUid
// taobao.openuid.get
//
// 获取授权账号对应的OpenUid
func TaobaoOpenuidGet(clt *core.SDKClient, req *util.TaobaoOpenuidGetAPIRequest, resp *util.TaobaoOpenuidGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
