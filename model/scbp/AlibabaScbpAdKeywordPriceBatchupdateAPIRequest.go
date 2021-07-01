package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpAdKeywordPriceBatchupdateAPIRequest
关键词批量改价 API请求
alibaba.scbp.ad.keyword.price.batchupdate

关键词批量改价 */
type AlibabaScbpAdKeywordPriceBatchupdateAPIRequest struct {
	model.Params
	// 系统自动生成
	_keywordUpdateDtoList []KeywordUpdateDto
}

// New
