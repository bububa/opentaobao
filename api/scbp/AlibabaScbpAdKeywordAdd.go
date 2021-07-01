package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

/* AlibabaScbpAdKeywordAdd
外贸直通车加词
alibaba.scbp.ad.keyword.add

外贸直通车加词服务 */
func AlibabaScbpAdKeywordAdd(clt *core.SDKClient, req *scbp.AlibabaScbpAdKeywordAddAPIRequest, session string) (*scbp.AlibabaScbpAdKeywordAddAPIResponse, error) {
	var resp scbp.AlibabaScbpAdKeywordAddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
