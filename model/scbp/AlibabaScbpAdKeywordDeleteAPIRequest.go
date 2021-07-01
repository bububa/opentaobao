package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpAdKeywordDeleteAPIRequest
外贸直通车删除关键词 API请求
alibaba.scbp.ad.keyword.delete

外贸直通车删除关键词 */
type AlibabaScbpAdKeywordDeleteAPIRequest struct {
	model.Params
	// 要删除的关键词
	_adKeyword string
}

// New
