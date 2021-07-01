package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

/* AlibabaScbpAdKeywordDeleteKeywordBatch
删除关键词
alibaba.scbp.ad.keyword.delete.keyword.batch

删除关键词 */
func AlibabaScbpAdKeywordDeleteKeywordBatch(clt *core.SDKClient, req *scbp.AlibabaScbpAdKeywordDeleteKeywordBatchAPIRequest, session string) (*scbp.AlibabaScbpAdKeywordDeleteKeywordBatchAPIResponse, error) {
	var resp scbp.AlibabaScbpAdKeywordDeleteKeywordBatchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
