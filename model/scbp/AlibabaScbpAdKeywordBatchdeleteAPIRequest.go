package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpAdKeywordBatchdeleteAPIRequest
外贸直通车批量删除关键词 API请求
alibaba.scbp.ad.keyword.batchdelete

外贸直通车批量删除关键词 */
type AlibabaScbpAdKeywordBatchdeleteAPIRequest struct {
	model.Params
	// 关键词Id列表
	_keywordIdList []int64
}

// New
