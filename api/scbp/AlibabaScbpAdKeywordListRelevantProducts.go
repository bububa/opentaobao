package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdKeywordListRelevantProducts 查询和词匹配的推广产品
// alibaba.scbp.ad.keyword.list.relevant.products
//
// 查询和词匹配的推广产品
func AlibabaScbpAdKeywordListRelevantProducts(clt *core.SDKClient, req *scbp.AlibabaScbpAdKeywordListRelevantProductsAPIRequest, session string) (*scbp.AlibabaScbpAdKeywordListRelevantProductsAPIResponse, error) {
	var resp scbp.AlibabaScbpAdKeywordListRelevantProductsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
