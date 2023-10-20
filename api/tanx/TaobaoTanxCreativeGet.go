package tanx

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tanx"
)

// TaobaoTanxCreativeGet 获取单个审核创意状态
// taobao.tanx.creative.get
//
// 获取单个审核创意状态
func TaobaoTanxCreativeGet(clt *core.SDKClient, req *tanx.TaobaoTanxCreativeGetAPIRequest, resp *tanx.TaobaoTanxCreativeGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
