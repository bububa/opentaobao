package aecreatives

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressAffiliateHotproductQueryAPIRequest
查询联盟爆品数据 API请求
aliexpress.affiliate.hotproduct.query

查询联盟爆品API */
type AliexpressAffiliateHotproductQueryAPIRequest struct {
	model.Params
	// 请求签名
	_appSignature string
	// 类目ID列表
	_categoryIds string
	// 返回字段列表
	_fields string
	// 关键词
	_keywords string
	// 最大售价
	_maxSalePrice int64
	// 最小售价
	_minSalePrice int64
	// 请求页数
	_pageNo int64
	// 每次请求数量
	_pageSize int64
	// 平台商家类型：ALL,PLAZA,TMALL
	_platformProductType string
	// 排序方式:SALE_PRICE_ASC, SALE_PRICE_DESC, DISCOUNT_ASC, DISCOUNT_DESC, LAST_VOLUME_ASC, LAST_VOLUME_DESC
	_sort string
	// 目标币种:USD, GBP, CAD, EUR, UAH, MXN, TRY, RUB, BRL, AUD, INR, JPY, IDR, SEK,KRW
	_targetCurrency string
	// 目标语言:EN,RU,PT,ES,FR,ID,IT,TH,JA,AR,VI,TR,DE,HE,KO,NL,PL,MX,CL,IW,IN
	_targetLanguage string
	// trackingId
	_trackingId string
	// 物流到达时间。3：3日达，5：5 日达，7：7日达，10：10日达
	_deliveryDays string
	// 商品收货国家，根据该国家税率政策返回对应商品价格
	_shipToCountry string
}

// New
