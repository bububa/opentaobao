package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpEffectKeywordSingleGet 单个关键词效果报表
// alibaba.scbp.effect.keyword.single.get
//
// 单个关键词效果报表
func AlibabaScbpEffectKeywordSingleGet(clt *core.SDKClient, req *scbp.AlibabaScbpEffectKeywordSingleGetAPIRequest, resp *scbp.AlibabaScbpEffectKeywordSingleGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
