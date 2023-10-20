package miniappopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniappopen"
)

// TaobaoMiniappDistributionMaterialGet 小程序投放---读取投放入口素材信息
// taobao.miniapp.distribution.material.get
//
// 读取已录入的投放入口素材信息。
func TaobaoMiniappDistributionMaterialGet(clt *core.SDKClient, req *miniappopen.TaobaoMiniappDistributionMaterialGetAPIRequest, resp *miniappopen.TaobaoMiniappDistributionMaterialGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
