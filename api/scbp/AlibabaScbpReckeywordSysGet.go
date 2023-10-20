package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpreckeywordsysget 系统推荐
// alibaba.scbp.reckeyword.sys.get
//
// 查询系统推荐词
func Alibabascbpreckeywordsysget(clt *core.SDKClient, req *scbp.AlibabascbpreckeywordsysgetAPIRequest, session string) (*scbp.AlibabascbpreckeywordsysgetAPIResponse, error) {
	var resp scbp.AlibabascbpreckeywordsysgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
