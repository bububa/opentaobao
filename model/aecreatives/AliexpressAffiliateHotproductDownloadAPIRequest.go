package aecreatives

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressAffiliateHotproductDownloadAPIRequest
联盟营销爆品下载接口 API请求
aliexpress.affiliate.hotproduct.download

查询联盟爆品API */
type AliexpressAffiliateHotproductDownloadAPIRequest struct {
	model.Params
	// trackingId
	_trackingId string
	// 请求签名
	_appSignature string
	// 类目ID
	_categoryId string
	// 返回字段列表
	_fields string
	// 站点商品标：global,it_site,es_site,ru_site
	_localeSite string
	// 请求页数
	_pageNo int64
	// 每次请求数量
	_pageSize int64
	// 目标币种:USD, GBP, CAD, EUR, UAH, MXN, TRY, RUB, BRL, AUD, INR, JPY, IDR, SEK,KRW
	_targetCurrency string
	// 目标语言:EN,RU,PT,ES,FR,ID,IT,TH,JA,AR,VI,TR,DE,HE,KO,NL,PL,MX,CL,IW,IN
	_targetLanguage string
	// 收货国家，可筛选能销售至该国家的商品，并根据该国家税率政策返回对应商品价格
	_country string
}

// New
