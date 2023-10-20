package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpreckeywordsearch 推荐词-词推词
// alibaba.scbp.reckeyword.search
//
// 推荐词-词推词
func Alibabascbpreckeywordsearch(clt *core.SDKClient, req *scbp.AlibabascbpreckeywordsearchAPIRequest, session string) (*scbp.AlibabascbpreckeywordsearchAPIResponse, error) {
	var resp scbp.AlibabascbpreckeywordsearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
