package travel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripTravelItemBaseModifyAPIRequest
【API3.0】度假线路商品编辑接口 API请求
taobao.alitrip.travel.item.base.modify

旅行度假新商品基本信息修改接口 第三版。提供商家通过TOP API方式修改商品除sku外的基本信息。 */
type TaobaoAlitripTravelItemBaseModifyAPIRequest struct {
	model.Params
	// 跟团游商品相关信息
	_groupItemExt *PontusTravelGroupItemExt
	// 商品基本信息
	_baseInfo *PontusTravelItemBaseInfo
	// 详细行程描述结构
	_itineraries []PontusTravelItemItineraryInfo
	// 商品id
	_itemId int64
	// 退款规则结构
	_refundInfo *PontusTravelItemRefundInfo
	// 预定规则结构
	_bookingRules []PontusTravelBookingRuleInfo
	// 自由行商品相关信息
	_freedomItemExt *PontusTravelFreedomItemExt
	// 商品的销售属性相关信息
	_salesInfo *PontusTravelItemSaleInfo
	// 可选，支持清空商品某些字段内容。多个待清空字段名之间以英文逗号分隔。目前支持清除以下字段内容：seller_cids表示删除关联的店铺类目id
	_fieldsToClean string
	// 邮轮商品相关扩展信息
	_cruiseItemExt *CruiseItemExt
	// 同城活动商品相关扩展信息
	_tcwlItemExt *TcwlItemExt
}

// New
