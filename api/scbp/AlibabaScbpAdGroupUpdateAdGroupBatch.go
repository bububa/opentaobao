package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

/* AlibabaScbpAdGroupUpdateAdGroupBatch
修改推广单元
alibaba.scbp.ad.group.update.ad.group.batch

修改推广单元 */
func AlibabaScbpAdGroupUpdateAdGroupBatch(clt *core.SDKClient, req *scbp.AlibabaScbpAdGroupUpdateAdGroupBatchAPIRequest, session string) (*scbp.AlibabaScbpAdGroupUpdateAdGroupBatchAPIResponse, error) {
	var resp scbp.AlibabaScbpAdGroupUpdateAdGroupBatchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
