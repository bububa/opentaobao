package travel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripTravelItemSkuPriceModifyAPIRequest
【API3.0】日期级别日历价格库存修改，增量维护 API请求
taobao.alitrip.travel.item.sku.price.modify

【API3.0】日期级别日历价格库存增量维护 */
type TaobaoAlitripTravelItemSkuPriceModifyAPIRequest struct {
	model.Params
	// 商品id。itemId和outProductId至少填写一个
	_itemId int64
	// 商品 外部商家编码。itemId和outProductId至少填写一个
	_outProductId string
	// 商品日历价格库存套餐
	_skus []PontusTravelItemSkuInfo
}

// New
