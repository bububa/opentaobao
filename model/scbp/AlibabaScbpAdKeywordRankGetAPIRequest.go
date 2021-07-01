package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpAdKeywordRankGetAPIRequest
获取外贸直通车关键词预估排名 API请求
alibaba.scbp.ad.keyword.rank.get

获取外贸直通车关键词预估排名 */
type AlibabaScbpAdKeywordRankGetAPIRequest struct {
	model.Params
	// 查询预估排名的关键词
	_keyword string
}

// New
