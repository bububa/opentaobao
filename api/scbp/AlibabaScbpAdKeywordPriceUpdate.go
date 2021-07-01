package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

/* AlibabaScbpAdKeywordPriceUpdate
关键词改价
alibaba.scbp.ad.keyword.price.update

关键词改价 */
func AlibabaScbpAdKeywordPriceUpdate(clt *core.SDKClient, req *scbp.AlibabaScbpAdKeywordPriceUpdateAPIRequest, session string) (*scbp.AlibabaScbpAdKeywordPriceUpdateAPIResponse, error) {
	var resp scbp.AlibabaScbpAdKeywordPriceUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
