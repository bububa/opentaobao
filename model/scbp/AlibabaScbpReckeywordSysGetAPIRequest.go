package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpReckeywordSysGetAPIRequest
系统推荐 API请求
alibaba.scbp.reckeyword.sys.get

查询系统推荐词 */
type AlibabaScbpReckeywordSysGetAPIRequest struct {
	model.Params
	// RecKeywordQuery
	_queryDto *RecKeywordQuery
}

// New
