package scbp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpEffectKeywordList 关键词报表
// alibaba.scbp.effect.keyword.list
//
// 关键词报表
func AlibabaScbpEffectKeywordList(ctx context.Context, clt *core.SDKClient, req *scbp.AlibabaScbpEffectKeywordListAPIRequest, resp *scbp.AlibabaScbpEffectKeywordListAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
