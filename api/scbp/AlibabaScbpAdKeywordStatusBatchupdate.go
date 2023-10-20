package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdKeywordStatusBatchupdate 批量启动暂停推广词状态
// alibaba.scbp.ad.keyword.status.batchupdate
//
// 批量启动暂停关键词推广状态
func AlibabaScbpAdKeywordStatusBatchupdate(clt *core.SDKClient, req *scbp.AlibabaScbpAdKeywordStatusBatchupdateAPIRequest, session string) (*scbp.AlibabaScbpAdKeywordStatusBatchupdateAPIResponse, error) {
	var resp scbp.AlibabaScbpAdKeywordStatusBatchupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
