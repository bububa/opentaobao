package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenShopSynchronizeAPIRequest
店铺同步接口 API请求
taobao.qimen.shop.synchronize

店铺同步接口描述 */
type TaobaoQimenShopSynchronizeAPIRequest struct {
	model.Params
	// 请求
	_request *TaobaoQimenShopSynchronizeRequest
}

// New
