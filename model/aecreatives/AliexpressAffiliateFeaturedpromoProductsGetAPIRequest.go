package aecreatives

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressAffiliateFeaturedpromoProductsGetAPIRequest
联盟主题推广活动商品信息获取 API请求
aliexpress.affiliate.featuredpromo.products.get

根据联盟主题推广活动或主题品库查询对应的商品。如下品库为固定品库，可长期调用。品库类型和名称如下：高佣品（Hot Product）、新品（New Arrival）、热销商品（Best Seller）、每周尖货（weeklydeals） */
type AliexpressAffiliateFeaturedpromoProductsGetAPIRequest struct {
	model.Params
	// 请求签名
	_appSignature string
	// 类目 ID 如何获取category_id，请参考，https://open.taobao.com/api.htm?docId=45801&docType=2&scopeId=17063
	_categoryId string
	// 返回字段列表
	_fields string
	// 查询页码
	_pageNo int64
	// 每页记录数，1-50
	_pageSize int64
	// 活动结束时间，PST 时区
	_promotionEndTime string
	// 主题活动的名称，如何获取主题活动，请参考"get featuredpromo info" API. 固定主题：高佣品（Hot Product）、新品（New Arrival）、热销商品（Best Seller）、每周尖货（weeklydeals）
	_promotionName string
	// 活动开始时间，PST 时区
	_promotionStartTime string
	// 排序方式：commissionAsc，commissionDesc, priceAsc，priceDesc，volumeAsc、volumeDesc, discountAsc, discountDesc, ratingAsc，ratingDesc, promotionTimeAsc, pr
	_sort string
	// 目标币种，可根据目标币种返回对应币种：USD, GBP, CAD, EUR, UAH, MXN, TRY, RUB, BRL, AUD, INR, JPY, IDR, SEK,KRW
	_targetCurrency string
	// 目标语言，可根据目标语言返回对应语言：EN,RU,PT,ES,FR,ID,IT,TH,JA,AR,VI,TR,DE,HE,KO,NL,PL,MX,CL,IN
	_targetLanguage string
	// trackingID
	_trackingId string
	// 商品收货国家，可筛选能销售至该国家的商品，并根据该国家税率政策返回对应商品价格
	_country string
}

// New
