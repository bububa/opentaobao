package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpadreportgetproductreport 产品报告
// alibaba.scbp.ad.report.get.product.report
//
// 产品报告
func Alibabascbpadreportgetproductreport(clt *core.SDKClient, req *scbp.AlibabascbpadreportgetproductreportAPIRequest, session string) (*scbp.AlibabascbpadreportgetproductreportAPIResponse, error) {
	var resp scbp.AlibabascbpadreportgetproductreportAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
