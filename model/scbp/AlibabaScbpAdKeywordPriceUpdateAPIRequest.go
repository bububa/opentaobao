package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpAdKeywordPriceUpdateAPIRequest
关键词改价 API请求
alibaba.scbp.ad.keyword.price.update

关键词改价 */
type AlibabaScbpAdKeywordPriceUpdateAPIRequest struct {
	model.Params
	// 只能取ascci字符
	_adKeyword string
	// 关键词价格单位元，一位小数
	_priceStr string
}

// New
