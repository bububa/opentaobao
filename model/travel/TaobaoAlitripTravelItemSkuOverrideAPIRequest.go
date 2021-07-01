package travel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripTravelItemSkuOverrideAPIRequest
【API3.0】商品级别日历价格库存修改，全量覆盖 API请求
taobao.alitrip.travel.item.sku.override

旅行度假新商品日历价格库存信息修改接口 第三版。提供商家通过TOP API方式修改商品sku信息。 */
type TaobaoAlitripTravelItemSkuOverrideAPIRequest struct {
	model.Params
	// 商品id。itemId和outProductId至少填写一个
	_itemId int64
	// 商品日历价格库存套餐
	_skus []PontusTravelItemSkuInfo
	// 商品 外部商家编码。itemId和outProductId至少填写一个
	_outProductId string
}

// New
