package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpkeywordmatchedproductsget 查询和词匹配的推广产品
// alibaba.scbp.keyword.matched.products.get
//
// 查询和词匹配的推广产品
func Alibabascbpkeywordmatchedproductsget(clt *core.SDKClient, req *scbp.AlibabascbpkeywordmatchedproductsgetAPIRequest, session string) (*scbp.AlibabascbpkeywordmatchedproductsgetAPIResponse, error) {
	var resp scbp.AlibabascbpkeywordmatchedproductsgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
