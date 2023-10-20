package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpadkeywordoperationpreferentialproduct 操作优推品
// alibaba.scbp.ad.keyword.operation.preferential.product
//
// 操作优推品
func Alibabascbpadkeywordoperationpreferentialproduct(clt *core.SDKClient, req *scbp.AlibabascbpadkeywordoperationpreferentialproductAPIRequest, session string) (*scbp.AlibabascbpadkeywordoperationpreferentialproductAPIResponse, error) {
	var resp scbp.AlibabascbpadkeywordoperationpreferentialproductAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
