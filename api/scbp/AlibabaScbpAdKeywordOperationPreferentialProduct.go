package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdKeywordOperationPreferentialProduct 操作优推品
// alibaba.scbp.ad.keyword.operation.preferential.product
//
// 操作优推品
func AlibabaScbpAdKeywordOperationPreferentialProduct(clt *core.SDKClient, req *scbp.AlibabaScbpAdKeywordOperationPreferentialProductAPIRequest, resp *scbp.AlibabaScbpAdKeywordOperationPreferentialProductAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
