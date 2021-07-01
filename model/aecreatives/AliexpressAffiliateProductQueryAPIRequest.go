package aecreatives

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressAffiliateProductQueryAPIRequest
联盟推广商品获取接口 API请求
aliexpress.affiliate.product.query

联盟推广商品搜索接口，用于搜索联盟推广商品数据 */
type AliexpressAffiliateProductQueryAPIRequest struct {
	model.Params
	// 安全签名
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
	// 查询页码
	_pageNo int64
	// 每页记录数
	_pageSize int64
	// 平台商品类型：ALL,PLAZA,TMALL
	_platformProductType string
	// 排序方式:SALE_PRICE_ASC, SALE_PRICE_DESC, LAST_VOLUME_ASC, LAST_VOLUME_DESC
	_sort string
	// 目标币种:USD, GBP, CAD, EUR, UAH, MXN, TRY, RUB, BRL, AUD, INR, JPY, IDR, SEK,KRW
	_targetCurrency string
	// 目标语言:EN,RU,PT,ES,FR,ID,IT,TH,JA,AR,VI,TR,DE,HE,KO,NL,PL,MX,CL,IW,IN
	_targetLanguage string
	// trackingId
	_trackingId string
	// 商品收货国家，根据该国家税率政策返回对应商品价格
	_shipToCountry string
	// 物流到达时间。3：3日达，5：5 日达，7：7日达，10：10日达
	_deliveryDays string
}

// New
