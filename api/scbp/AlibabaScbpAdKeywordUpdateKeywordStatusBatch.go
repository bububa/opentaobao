package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdKeywordUpdateKeywordStatusBatch 修改关键词状态
// alibaba.scbp.ad.keyword.update.keyword.status.batch
//
// 修改关键词状态
func AlibabaScbpAdKeywordUpdateKeywordStatusBatch(clt *core.SDKClient, req *scbp.AlibabaScbpAdKeywordUpdateKeywordStatusBatchAPIRequest, session string) (*scbp.AlibabaScbpAdKeywordUpdateKeywordStatusBatchAPIResponse, error) {
	var resp scbp.AlibabaScbpAdKeywordUpdateKeywordStatusBatchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
