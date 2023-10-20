package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdKeywordPriceBatchupdate 关键词批量改价
// alibaba.scbp.ad.keyword.price.batchupdate
//
// 关键词批量改价
func AlibabaScbpAdKeywordPriceBatchupdate(clt *core.SDKClient, req *scbp.AlibabaScbpAdKeywordPriceBatchupdateAPIRequest, resp *scbp.AlibabaScbpAdKeywordPriceBatchupdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
