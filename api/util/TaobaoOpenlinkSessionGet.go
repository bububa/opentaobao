package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// TaobaoOpenlinkSessionGet 获取授权session信息
// taobao.openlink.session.get
//
// 帮助第三方isv生成三方session
func TaobaoOpenlinkSessionGet(clt *core.SDKClient, req *util.TaobaoOpenlinkSessionGetAPIRequest, resp *util.TaobaoOpenlinkSessionGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
