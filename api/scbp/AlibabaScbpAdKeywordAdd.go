package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdKeywordAdd 外贸直通车加词
// alibaba.scbp.ad.keyword.add
//
// 外贸直通车加词服务
func AlibabaScbpAdKeywordAdd(clt *core.SDKClient, req *scbp.AlibabaScbpAdKeywordAddAPIRequest, resp *scbp.AlibabaScbpAdKeywordAddAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
