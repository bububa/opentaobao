package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdKeywordDelete 外贸直通车删除关键词
// alibaba.scbp.ad.keyword.delete
//
// 外贸直通车删除关键词
func AlibabaScbpAdKeywordDelete(clt *core.SDKClient, req *scbp.AlibabaScbpAdKeywordDeleteAPIRequest, session string) (*scbp.AlibabaScbpAdKeywordDeleteAPIResponse, error) {
	var resp scbp.AlibabaScbpAdKeywordDeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
