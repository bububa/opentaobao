package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdKeywordOperationPreferentialProduct 操作优推品
// alibaba.scbp.ad.keyword.operation.preferential.product
//
// 操作优推品
func AlibabaScbpAdKeywordOperationPreferentialProduct(clt *core.SDKClient, req *scbp.AlibabaScbpAdKeywordOperationPreferentialProductAPIRequest, session string) (*scbp.AlibabaScbpAdKeywordOperationPreferentialProductAPIResponse, error) {
	var resp scbp.AlibabaScbpAdKeywordOperationPreferentialProductAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
