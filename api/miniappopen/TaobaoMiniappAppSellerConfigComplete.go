package miniappopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniappopen"
)

// TaobaoMiniappAppSellerConfigComplete 商家完成小程序相关配置
// taobao.miniapp.app.seller.config.complete
//
// 通过该接口告知平台商家已经完成小程序相关的必要设置，可进行后续操作。主要用于小部件、客服插件等场景。
func TaobaoMiniappAppSellerConfigComplete(clt *core.SDKClient, req *miniappopen.TaobaoMiniappAppSellerConfigCompleteAPIRequest, resp *miniappopen.TaobaoMiniappAppSellerConfigCompleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
