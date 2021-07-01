package aecreatives

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressAffiliateProductdetailGetAPIRequest
联盟商品详情获取接口 API请求
aliexpress.affiliate.productdetail.get

联盟推广商品搜索接口，用于搜索联盟推广商品数据 */
type AliexpressAffiliateProductdetailGetAPIRequest struct {
	model.Params
	// 安全签名
	_appSignature string
	// commission_rate,sale_price
	_fields string
	// 商品ID列表
	_productIds string
	// 目标币种:USD, GBP, CAD, EUR, UAH, MXN, TRY, RUB, BRL, AUD, INR, JPY, IDR, SEK,KRW
	_targetCurrency string
	// 目标语言:EN,RU,PT,ES,FR,ID,IT,TH,JA,AR,VI,TR,DE,HE,KO,NL,PL,MX,CL,IW,IN
	_targetLanguage string
	// trackingId
	_trackingId string
	// 商品收货国家，可筛选能销售至该国家的商品，并根据该国家税率政策返回对应商品价格
	_country string
}

// New
