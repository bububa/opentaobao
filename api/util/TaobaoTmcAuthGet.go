package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// TaobaoTmcAuthGet TMC授权token
// taobao.tmc.auth.get
//
// TMC连接授权Token
func TaobaoTmcAuthGet(clt *core.SDKClient, req *util.TaobaoTmcAuthGetAPIRequest, resp *util.TaobaoTmcAuthGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
