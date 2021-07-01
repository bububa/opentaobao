package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFuwuSkuGetAPIRequest
获取内购服务及SKU详情 API请求
taobao.fuwu.sku.get

通过服务code和用户nick，获取该服务对应的收费项目的sku信息，包括价格、可购买周期、用户能否购买等信息 */
type TaobaoFuwuSkuGetAPIRequest struct {
	model.Params
	// 服务code
	_articleCode string
	// 用户的淘宝nick
	_nick string
}

// New
