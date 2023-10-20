package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpadcustomerfindcustomerinfo 查询客户信息
// alibaba.scbp.ad.customer.find.customer.info
//
// 查询客户信息
func Alibabascbpadcustomerfindcustomerinfo(clt *core.SDKClient, req *scbp.AlibabascbpadcustomerfindcustomerinfoAPIRequest, session string) (*scbp.AlibabascbpadcustomerfindcustomerinfoAPIResponse, error) {
	var resp scbp.AlibabascbpadcustomerfindcustomerinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
