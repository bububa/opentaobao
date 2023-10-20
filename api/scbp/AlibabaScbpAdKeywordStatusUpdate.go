package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdKeywordStatusUpdate 关键词启动暂停推广
// alibaba.scbp.ad.keyword.status.update
//
// 关键词启动暂停推广
func AlibabaScbpAdKeywordStatusUpdate(clt *core.SDKClient, req *scbp.AlibabaScbpAdKeywordStatusUpdateAPIRequest, resp *scbp.AlibabaScbpAdKeywordStatusUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
