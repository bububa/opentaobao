package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdCustomerFindCustomerInfo 查询客户信息
// alibaba.scbp.ad.customer.find.customer.info
//
// 查询客户信息
func AlibabaScbpAdCustomerFindCustomerInfo(clt *core.SDKClient, req *scbp.AlibabaScbpAdCustomerFindCustomerInfoAPIRequest, session string) (*scbp.AlibabaScbpAdCustomerFindCustomerInfoAPIResponse, error) {
	var resp scbp.AlibabaScbpAdCustomerFindCustomerInfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
