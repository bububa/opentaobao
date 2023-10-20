package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpeffectproductreport 所有产品报表
// alibaba.scbp.effect.product.report
//
// 所有产品报表
func Alibabascbpeffectproductreport(clt *core.SDKClient, req *scbp.AlibabascbpeffectproductreportAPIRequest, session string) (*scbp.AlibabascbpeffectproductreportAPIResponse, error) {
	var resp scbp.AlibabascbpeffectproductreportAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
