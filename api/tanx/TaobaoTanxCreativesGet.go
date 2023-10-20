package tanx

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tanx"
)

// TaobaoTanxCreativesGet 批量获取DSP用户的创意审核结果
// taobao.tanx.creatives.get
//
// 批量获取DSP用户的创意审核结果
func TaobaoTanxCreativesGet(clt *core.SDKClient, req *tanx.TaobaoTanxCreativesGetAPIRequest, resp *tanx.TaobaoTanxCreativesGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
