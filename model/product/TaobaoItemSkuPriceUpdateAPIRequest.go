package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoItemSkuPriceUpdateAPIRequest
更新商品SKU的价格 API请求
taobao.item.sku.price.update

更新商品SKU的价格 */
type TaobaoItemSkuPriceUpdateAPIRequest struct {
	model.Params
	// Sku所属商品数字id，可通过 taobao.item.get 获取
	_numIid int64
	// Sku属性串。格式:pid:vid;pid:vid,如: 1627207:3232483;1630696:3284570,表示机身颜色:军绿色;手机套餐:一电一充
	_properties string
	// Sku的库存数量。sku的总数量应该小于等于商品总数量(Item的NUM)，sku数量变化后item的总数量也会随着变化。取值范围:大于等于零的整数
	_quantity int64
	// Sku的销售价格。精确到2位小数;单位:元。如:200.07，表示:200元7分。修改后的sku价格要保证商品的价格在所有sku价格所形成的价格区间内（例如：商品价格为6元，sku价格有5元、10元两种，如果要修改5元sku的价格，那么修改的范围只能是0-6元之间；如果要修改10元的sku，那么修改的范围只能是6到无穷大的区间中）
	_price float64
	// Sku的商家外部id
	_outerId string
	// sku所属商品的价格。当用户更新sku，使商品价格不属于sku价格之间的时候，用于修改商品的价格，使sku能够更新成功
	_itemPrice float64
	// Sku文字的版本。可选值:zh_HK(繁体),zh_CN(简体);默认值:zh_CN
	_lang string
	// 忽略警告提示.
	_ignorewarning string
}

// New
