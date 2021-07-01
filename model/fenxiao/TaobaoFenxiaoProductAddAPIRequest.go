package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFenxiaoProductAddAPIRequest
添加产品 API请求
taobao.fenxiao.product.add

添加分销平台产品数据。业务逻辑与分销系统前台页面一致。<br/><br/>    * 产品图片默认为空<br/>    * 产品发布后默认为下架状态 */
type TaobaoFenxiaoProductAddAPIRequest struct {
	model.Params
	// 运费类型，可选值：seller（供应商承担运费）、buyer（分销商承担运费）,默认seller。
	_postageType string
	// 是否有发票，可选值：false（否）、true（是），默认false。
	_haveInvoice string
	// 是否有保修，可选值：false（否）、true（是），默认false。
	_haveQuarantee string
	// 分销方式：AGENT（只做代销，默认值）、DEALER（只做经销）、ALL（代销和经销都做）
	_tradeType string
	// 产品名称，长度不超过60个字节。
	_name string
	// 产品线ID
	_productcatId int64
	// 采购基准价格，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
	_standardPrice string
	// 代销采购价格，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
	_costPrice string
	// 最低零售价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
	_retailPriceLow string
	// 最高零售价，单位：元。例：“10.56”。必须在0.01元到10000000元之间，最高零售价必须大于最低零售价。
	_retailPriceHigh string
	// 所属类目id，参考Taobao.itemcats.get，不支持成人等类目，输入成人类目id保存提示类目属性错误。
	_categoryId int64
	// 商家编码，长度不能超过60个字节。
	_outerId string
	// 产品库存必须是1到999999。
	_quantity int64
	// 产品描述，长度为5到25000字符。
	_desc string
	// 所在地：省，例：“浙江”
	_prov string
	// 所在地：市，例：“杭州”
	_city string
	// 运费模板ID，参考taobao.postages.get。
	_postageId int64
	// 平邮费用，单位：元。例：“10.56”。 大小为0.01元到999999元之间。
	_postageOrdinary string
	// 快递费用，单位：元。例：“10.56”。 大小为0.01元到999999元之间。
	_postageFast string
	// ems费用，单位：元。例：“10.56”。 大小为0.00元到999999元之间。
	_postageEms string
	// 折扣ID
	_discountId int64
	// 添加产品时，添加入参isAuthz:yes|no <br/>yes:需要授权 <br/>no:不需要授权 <br/>默认是需要授权
	_isAuthz string
	// 产品主图图片空间相对路径或绝对路径
	_picPath string
	// 产品主图，大小不超过500k，格式为gif,jpg,jpeg,png,bmp等图片
	_image *model.File
	// 产品属性，格式为pid:vid;pid:vid
	_properties string
	// 属性别名，格式为：pid:vid:alias;pid:vid:alias（alias为别名）
	_propertyAlias string
	// 自定义属性。格式为pid:value;pid:value
	_inputProperties string
	// sku的采购基准价。如果多个，用逗号分隔，并与其他sku信息保持相同顺序
	_skuStandardPrices string
	// sku的采购价。如果多个，用逗号分隔，并与其他sku信息保持相同顺序
	_skuCostPrices string
	// sku的商家编码。如果多个，用逗号分隔，并与其他sku信息保持相同顺序
	_skuOuterIds string
	// sku的库存。如果多个，用逗号分隔，并与其他sku信息保持相同顺序
	_skuQuantitys string
	// sku的属性。如果多个，用逗号分隔，并与其他sku信息保持相同顺序
	_skuProperties string
	// 经销采购价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
	_dealerCostPrice string
	// sku的经销采购价。如果多个，用逗号分隔，并与其他sku信息保持相同顺序。其中每个值的单位：元。例：“10.56,12.3”。必须在0.01元到10000000元之间。
	_skuDealerCostPrices string
	// 导入的商品ID
	_itemId int64
	// 零售基准价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
	_standardRetailPrice string
	// 产品spuID，达尔文产品必须要传spuID，否则不能发布。其他非达尔文产品，看情况传
	_spuId int64
}

// New
