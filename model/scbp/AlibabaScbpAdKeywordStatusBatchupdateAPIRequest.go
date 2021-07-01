package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpAdKeywordStatusBatchupdateAPIRequest
批量启动暂停推广词状态 API请求
alibaba.scbp.ad.keyword.status.batchupdate

批量启动暂停关键词推广状态 */
type AlibabaScbpAdKeywordStatusBatchupdateAPIRequest struct {
	model.Params
	// 系统自动生成
	_keywordUpdateDtoList []KeywordUpdateDto
}

// New
