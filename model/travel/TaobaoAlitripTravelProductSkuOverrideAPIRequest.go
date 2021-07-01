package travel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripTravelProductSkuOverrideAPIRequest
（供销）产品级别日历价格库存修改，全量覆盖 API请求
taobao.alitrip.travel.product.sku.override

（供销）产品级别日历价格库存修改，全量覆盖 */
type TaobaoAlitripTravelProductSkuOverrideAPIRequest struct {
	model.Params
	// 商品 外部商家编码。itemId和outProductId至少填写一个
	_outProductId string
	// 商品id。itemId和outProductId至少填写一个
	_itemId int64
	// 商品日历价格库存套餐
	_skus []ItemSkuInfo
}

// New
