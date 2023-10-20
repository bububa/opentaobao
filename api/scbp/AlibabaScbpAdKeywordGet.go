package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdKeywordGet 外贸直通车查询关键词
// alibaba.scbp.ad.keyword.get
//
// 外贸直通车查询关键词
func AlibabaScbpAdKeywordGet(clt *core.SDKClient, req *scbp.AlibabaScbpAdKeywordGetAPIRequest, resp *scbp.AlibabaScbpAdKeywordGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
