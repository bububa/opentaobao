package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdGroupCreateAdGroupBatch 创建推广单元
// alibaba.scbp.ad.group.create.ad.group.batch
//
// 创建推广单元
func AlibabaScbpAdGroupCreateAdGroupBatch(clt *core.SDKClient, req *scbp.AlibabaScbpAdGroupCreateAdGroupBatchAPIRequest, resp *scbp.AlibabaScbpAdGroupCreateAdGroupBatchAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
