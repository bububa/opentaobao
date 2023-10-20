package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpproductstatusupdate 修改P4P产品推广状态
// alibaba.scbp.product.status.update
//
// 修改P4P产品推广状态
func Alibabascbpproductstatusupdate(clt *core.SDKClient, req *scbp.AlibabascbpproductstatusupdateAPIRequest, session string) (*scbp.AlibabascbpproductstatusupdateAPIResponse, error) {
	var resp scbp.AlibabascbpproductstatusupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
