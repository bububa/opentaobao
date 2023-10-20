package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdKeywordPriceBatchupdate 关键词批量改价
// alibaba.scbp.ad.keyword.price.batchupdate
//
// 关键词批量改价
func AlibabaScbpAdKeywordPriceBatchupdate(clt *core.SDKClient, req *scbp.AlibabaScbpAdKeywordPriceBatchupdateAPIRequest, session string) (*scbp.AlibabaScbpAdKeywordPriceBatchupdateAPIResponse, error) {
	var resp scbp.AlibabaScbpAdKeywordPriceBatchupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
