package uscesl

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoUsceslIteminfoPutAPIRequest
电子价签显示用商品信息写入 API请求
taobao.uscesl.iteminfo.put

用于电子价签上显示的商品信息的写入，包含价格及促销信息 */
type TaobaoUsceslIteminfoPutAPIRequest struct {
	model.Params
	// 型号
	_modelNum string
	// 价格单位
	_priceUnit string
	// 品牌名
	_brandName string
	// 销售规格
	_saleSpec string
	// 分类
	_categoryName string
	// 等级
	_rank string
	// 商品变更状态
	_itemChangeStatus string
	// 实际销售价格，单位：分
	_acctionPrice string
	// 能效
	_energyEfficiency string
	// 商品skuId
	_skuId string
	// 促销开始时间:yyyy-MM-dd HH:mm:ss
	_promotionStart string
	// 商品条码
	_itemBarCode string
	// 商品名称
	_itemTitle string
	// 促销文案
	_promotionText string
	// 扩展属性C
	_customizeFeatureC string
	// 扩展属性D
	_customizeFeatureD string
	// 扩展属性E
	_customizeFeatureE string
	// 扩展属性F
	_customizeFeatureF string
	// 扩展属性G
	_customizeFeatureG string
	// 扩展属性H
	_customizeFeatureH string
	// 扩展属性I
	_customizeFeatureI string
	// 扩展属性J
	_customizeFeatureJ string
	// 二维码
	_itemQrCode string
	// 商品状态0:在售 1:售罄
	_itemStatus int64
	// 促销状态0:非促销 1:促销
	_ifPromotion bool
	// 商品编码
	_itemId int64
	// 促销结束时间:yyyy-MM-dd HH:mm:ss
	_promotionEnd string
	// 促销原因
	_promotionReason string
	// 商品原价
	_originalPrice string
	// 商品简称
	_shortTitle string
	// 扩展属性B
	_customizeFeatureB string
	// 产地
	_productionPlace string
	// 扩展属性A
	_customizeFeatureA string
	// 门店ID
	_shopId int64
}

// New
