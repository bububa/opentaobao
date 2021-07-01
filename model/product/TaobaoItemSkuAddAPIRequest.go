package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoItemSkuAddAPIRequest
添加SKU API请求
taobao.item.sku.add

新增一个sku到num_iid指定的商品中 <br/>传入的iid所对应的商品必须属于当前会话的用户 */
type TaobaoItemSkuAddAPIRequest struct {
	model.Params
	// Sku所属商品数字id。必选
	_numIid int64
	// Sku属性串。格式:pid:vid;pid:vid,如:1627207:3232483;1630696:3284570,表示:机身颜色:军绿色;手机套餐:一电一充。
	_properties string
	// Sku的库存数量。sku的总数量应该小于等于商品总数量(Item的NUM)。取值范围:大于零的整数
	_quantity int64
	// Sku的销售价格。商品的价格要在商品所有的sku的价格之间。精确到2位小数;单位:元。如:200.07，表示:200元7分
	_price float64
	// Sku的商家外部id
	_outerId string
	// sku所属商品的价格。当用户新增sku，使商品价格不属于sku价格之间的时候，用于修改商品的价格，使sku能够添加成功
	_itemPrice float64
	// Sku文字的版本。可选值:zh_HK(繁体),zh_CN(简体);默认值:zh_CN
	_lang string
	// 忽略警告提示.
	_ignorewarning string
}

// New
