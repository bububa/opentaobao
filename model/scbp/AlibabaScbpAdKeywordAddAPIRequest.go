package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpAdKeywordAddAPIRequest
外贸直通车加词 API请求
alibaba.scbp.ad.keyword.add

外贸直通车加词服务 */
type AlibabaScbpAdKeywordAddAPIRequest struct {
	model.Params
	// 加入的词
	_adKeyword string
	// 词的出价
	_priceStr string
	// 分组名
	_tagName string
}

// New
