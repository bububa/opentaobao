package wlb

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWlbItemAddAPIRequest
添加单个物流宝商品 API请求
taobao.wlb.item.add

添加物流宝商品，支持物流宝子商品和属性添加 */
type TaobaoWlbItemAddAPIRequest struct {
	model.Params
	// 商品名称
	_name string
	// 商品标题
	_title string
	// 商品编码
	_itemCode string
	// 商品备注
	_remark string
	// NORMAL--普通商品<br/>COMBINE--组合商品<br/>DISTRIBUTION--分销
	_type string
	// 是否sku
	_isSku string
	// 属性名列表,目前支持的属性：<br/>毛重:GWeight	<br/>净重:Nweight<br/>皮重: Tweight<br/>自定义属性：<br/>paramkey1<br/>paramkey2<br/>paramkey3<br/>paramkey4
	_proNameList string
	// 属性值列表：<br/>10,8
	_proValueList string
	// 是否易碎品
	_isFriable bool
	// 是否危险品
	_isDangerous bool
	// 商品颜色
	_color string
	// 商品重量，单位G
	_weight int64
	// 商品长度，单位毫米
	_length int64
	// 商品宽度，单位毫米
	_width int64
	// 商品高度，单位毫米
	_height int64
	// 商品体积，单位立方厘米
	_volume int64
	// 货类
	_goodsCat string
	// 计价货类
	_pricingCat string
	// 商品包装材料类型
	_packageMaterial string
	// 商品价格，单位：分
	_price int64
	// 是否支持批次
	_supportBatch bool
}

// New
