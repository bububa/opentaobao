package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdKeywordPriceUpdate 关键词改价
// alibaba.scbp.ad.keyword.price.update
//
// 关键词改价
func AlibabaScbpAdKeywordPriceUpdate(clt *core.SDKClient, req *scbp.AlibabaScbpAdKeywordPriceUpdateAPIRequest, resp *scbp.AlibabaScbpAdKeywordPriceUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
