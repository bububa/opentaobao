package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpProductList 查询P4P产品
// alibaba.scbp.product.list
//
// 查询P4P产品
func AlibabaScbpProductList(clt *core.SDKClient, req *scbp.AlibabaScbpProductListAPIRequest, resp *scbp.AlibabaScbpProductListAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
