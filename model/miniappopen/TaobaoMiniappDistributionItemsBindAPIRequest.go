package miniappopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoMiniappDistributionItemsBindAPIRequest
小程序投放-商品绑定/解绑 API请求
taobao.miniapp.distribution.items.bind

提供给使用了投放插件的服务商，可以调用该API实现帮助商家更新已创建的投放单中的绑定商品信息。 */
type TaobaoMiniappDistributionItemsBindAPIRequest struct {
	model.Params
	// 商品id列表
	_targetEntityList []string
	// 投放的商家应用完整链接
	_url string
	// true表示新增绑定，false表示解绑
	_addBind bool
}

// New
