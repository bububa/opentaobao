package aecreatives

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressAffiliateImageSearchAPIRequest
图搜 API请求
aliexpress.affiliate.image.search

图片搜索接口 */
type AliexpressAffiliateImageSearchAPIRequest struct {
	model.Params
	// API signature
	_appSignature string
	// 请求字段
	_fields string
	// 图片文件字节数组，最大不超过 100 KB
	_imageFileBytes *model.File
	// 图片类目倾向，不填则为最佳匹配。0 - 服装；3 - 包；4 - 鞋子；88888888 - 其他类目
	_imgCid string
	// 媒体用户唯一识别号
	_mediaUserId string
	// 搜索结果数量，最高 50
	_productCnt int64
	// ship-to 国家
	_shptTo string
	// 排序方式:SALE_PRICE_ASC, SALE_PRICE_DESC,LAST_VOLUME_ASC, LAST_VOLUME_DESC
	_sort string
	// 目标币种:USD, GBP, CAD, EUR, UAH, MXN,TRY, RUB, BRL, AUD, INR, JPY, IDR, SEK,KRW
	_targetCurrency string
	// 目标语言:en,ru,pt,es,fr,id,it,th,ja,ar,vi,tr,de,he,ko,nl,pl,mx,cl,iw,in
	_targetLanguage string
	// 媒体 trackingid
	_trackingId string
}

// New
