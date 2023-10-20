package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdKeywordCreateKeywordBatch 关键词添加
// alibaba.scbp.ad.keyword.create.keyword.batch
//
// 关键词添加
func AlibabaScbpAdKeywordCreateKeywordBatch(clt *core.SDKClient, req *scbp.AlibabaScbpAdKeywordCreateKeywordBatchAPIRequest, session string) (*scbp.AlibabaScbpAdKeywordCreateKeywordBatchAPIResponse, error) {
	var resp scbp.AlibabaScbpAdKeywordCreateKeywordBatchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
