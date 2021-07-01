package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFenxiaoProductUpdateAPIRequest
更新产品 API请求
taobao.fenxiao.product.update

更新分销平台产品数据，不传更新数据返回失败<br><br/>1. 对sku进行增、删操作时，原有的sku_ids字段会被忽略，请使用sku_properties和sku_properties_del。<br> */
type TaobaoFenxiaoProductUpdateAPIRequest struct {
	model.Params
	// 运费类型，可选值：seller（供应商承担运费）、buyer（分销商承担运费）。
	_postageType string
	// 是否有发票，可选值：false（否）、true（是），默认false。
	_haveInvoice string
	// 是否有保修，可选值：false（否）、true（是），默认false。
	_haveQuarantee string
	// 发布状态，可选值：up（上架）、down（下架）、delete（删除），输入非法字符则忽略。
	_status string
	// 产品ID
	_pid int64
	// 产品名称，长度不超过60个字节。
	_name string
	// 采购基准价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
	_standardPrice string
	// 代销采购价格，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
	_costPrice string
	// 最低零售价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
	_retailPriceLow string
	// 最高零售价，单位：元。例：“10.56”。必须在0.01元到10000000元之间，最高零售价必须大于最低零售价。
	_retailPriceHigh string
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
	// 运费模板ID，参考taobao.postages.get。更新时必须指定运费类型为 buyer，否则不更新。
	_postageId int64
	// 平邮费用，单位：元。例：“10.56”。大小为0.01元到999999元之间。更新时必须指定运费类型为buyer，否则不更新。
	_postageOrdinary string
	// 快递费用，单位：元。例：“10.56”。大小为0.01元到999999元之间。更新时必须指定运费类型为buyer，否则不更新。
	_postageFast string
	// ems费用，单位：元。例：“10.56”。大小为0.01元到999999元之间。更新时必须指定运费类型为buyer，否则不更新。
	_postageEms string
	// sku id列表，例：1001,1002,1003。如果传入sku_properties将忽略此参数。
	_skuIds string
	// sku采购价格，单位元，例："10.50,11.00,20.50"，字段必须和上面的sku_ids或sku_properties保持一致。
	_skuCostPrices string
	// sku库存，单位元，例："10,20,30"，字段必须和sku_ids或sku_properties保持一致。
	_skuQuantitys string
	// sku商家编码 ，单位元，例："S1000,S1002,S1003"，字段必须和上面的id或sku_properties保持一致，如果没有可以写成",,"
	_skuOuterIds string
	// 折扣ID
	_discountId int64
	// sku采购基准价，单位元，例："10.50,11.00,20.50"，字段必须和上面的sku_ids或sku_properties保持一致。
	_skuStandardPrices string
	// sku属性。格式:pid:vid;pid:vid,表示一组属性如:1627207:3232483;1630696:3284570,表示一组:机身颜色:军绿色;手机套餐:一电一充。多组之间用逗号“,”区分。(属性的pid调用taobao.itemprops.get取得，属性值的vid用taobao.itempropvalues.get取得vid)<br/>通过此字段可新增和更新sku。若传入此值将忽略sku_ids字段。sku其他字段与此值保持一致。
	_skuProperties string
	// 根据sku属性删除sku信息。需要按组删除属性。
	_skuPropertiesDel string
	// 产品是否需要授权isAuthz:yes|no <br/>yes:需要授权 <br/>no:不需要授权
	_isAuthz string
	// 产品主图图片空间相对路径或绝对路径
	_picPath string
	// 主图图片，如果pic_path参数不空，则优先使用pic_path，忽略该参数
	_image *model.File
	// 产品属性
	_properties string
	// 属性别名
	_propertyAlias string
	// 自定义属性。格式为pid:value;pid:value
	_inputProperties string
	// 经销采购价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
	_dealerCostPrice string
	// sku的经销采购价。如果多个，用逗号分隔，并与其他sku信息保持相同顺序。其中每个值的单位：元。例：“10.56,12.3”。必须在0.01元到10000000元之间。
	_skuDealerCostPrices string
	// 所属类目id，参考Taobao.itemcats.get，不支持成人等类目，输入成人类目id保存提示类目属性错误。
	_categoryId int64
	// 零售基准价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
	_standardRetailPrice string
}

// New
