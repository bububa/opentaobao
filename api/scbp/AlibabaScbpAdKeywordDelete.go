package scbp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdKeywordDelete 外贸直通车删除关键词
// alibaba.scbp.ad.keyword.delete
//
// 外贸直通车删除关键词
func AlibabaScbpAdKeywordDelete(ctx context.Context, clt *core.SDKClient, req *scbp.AlibabaScbpAdKeywordDeleteAPIRequest, resp *scbp.AlibabaScbpAdKeywordDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
