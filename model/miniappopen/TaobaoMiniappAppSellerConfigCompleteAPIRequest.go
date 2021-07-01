package miniappopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoMiniappAppSellerConfigCompleteAPIRequest
商家完成小程序相关配置 API请求
taobao.miniapp.app.seller.config.complete

通过该接口告知平台商家已经完成小程序相关的必要设置，可进行后续操作。主要用于小部件、客服插件等场景。 */
type TaobaoMiniappAppSellerConfigCompleteAPIRequest struct {
	model.Params
	// 商家已完成配置的小部件/B端插件的appid
	_appId int64
	// 小部件必传，B端插件不用传。与app_id对应的已完成配置的版本号
	_version string
}

// New
