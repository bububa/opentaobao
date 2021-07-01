package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

/* AlibabaScbpAdKeywordGetKeywordCountByQuery
计划关键词数目
alibaba.scbp.ad.keyword.get.keyword.count.by.query

计划关键词数目 */
func AlibabaScbpAdKeywordGetKeywordCountByQuery(clt *core.SDKClient, req *scbp.AlibabaScbpAdKeywordGetKeywordCountByQueryAPIRequest, session string) (*scbp.AlibabaScbpAdKeywordGetKeywordCountByQueryAPIResponse, error) {
	var resp scbp.AlibabaScbpAdKeywordGetKeywordCountByQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
