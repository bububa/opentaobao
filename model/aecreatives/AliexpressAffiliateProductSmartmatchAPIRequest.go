package aecreatives

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressAffiliateProductSmartmatchAPIRequest
联盟物料智能推荐api API请求
aliexpress.affiliate.product.smartmatch

联盟物料算法智能推荐 */
type AliexpressAffiliateProductSmartmatchAPIRequest struct {
	model.Params
	// 接入APP信息
	_app string
	// 请求签名
	_appSignature string
	// 设备信息
	_device string
	// adid或者idfa
	_deviceId string
	// 返回字段列表
	_fields string
	// 关键词
	_keywords string
	// 商品ID
	_productId string
	// 站点信息
	_site string
	// 目标币种:USD, GBP, CAD, EUR, UAH, MXN, TRY, RUB, BRL, AUD, INR, JPY, IDR, SEK,KRW
	_targetCurrency string
	// 目标语言:EN,RU,PT,ES,FR,ID,IT,TH,JA,AR,VI,TR,DE,HE,KO,NL,PL,MX,CL,IW,IN
	_targetLanguage string
	// trackingId
	_trackingId string
	// 用户信息
	_user string
	// 请求页数
	_pageNo int64
	// 商品收货国家，可筛选能销售至该国家的商品，并根据该国家税率政策返回对应商品价格
	_country string
}

// New
