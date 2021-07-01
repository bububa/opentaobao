package travel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripTravelItemSkuPackageModifyAPIRequest
【API3.0】套餐级别日历价格库存增删操作 API请求
taobao.alitrip.travel.item.sku.package.modify

【API3.0】套餐级别日历价格库存增删操作 */
type TaobaoAlitripTravelItemSkuPackageModifyAPIRequest struct {
	model.Params
	// 商品id。itemId和outProductId至少填写一个
	_itemId int64
	// 商品 外部商家编码。itemId和outProductId至少填写一个
	_outProductId string
	// 商品日历价格库存套餐
	_skus []ItemSkuInfo
}

// New
