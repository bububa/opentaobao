package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpadkeywordlistrelevantproducts 查询和词匹配的推广产品
// alibaba.scbp.ad.keyword.list.relevant.products
//
// 查询和词匹配的推广产品
func Alibabascbpadkeywordlistrelevantproducts(clt *core.SDKClient, req *scbp.AlibabascbpadkeywordlistrelevantproductsAPIRequest, session string) (*scbp.AlibabascbpadkeywordlistrelevantproductsAPIResponse, error) {
	var resp scbp.AlibabascbpadkeywordlistrelevantproductsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
