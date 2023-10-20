package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdKeywordBatchQueryKeywordRankPrice 批量查询关键词前五名排价
// alibaba.scbp.ad.keyword.batch.query.keyword.rank.price
//
// 批量查询关键词前五名排价
func AlibabaScbpAdKeywordBatchQueryKeywordRankPrice(clt *core.SDKClient, req *scbp.AlibabaScbpAdKeywordBatchQueryKeywordRankPriceAPIRequest, session string) (*scbp.AlibabaScbpAdKeywordBatchQueryKeywordRankPriceAPIResponse, error) {
	var resp scbp.AlibabaScbpAdKeywordBatchQueryKeywordRankPriceAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
