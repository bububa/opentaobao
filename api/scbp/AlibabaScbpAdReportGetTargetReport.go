package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpadreportgettargetreport 定向报告
// alibaba.scbp.ad.report.get.target.report
//
// 定向报告
func Alibabascbpadreportgettargetreport(clt *core.SDKClient, req *scbp.AlibabascbpadreportgettargetreportAPIRequest, session string) (*scbp.AlibabascbpadreportgettargetreportAPIResponse, error) {
	var resp scbp.AlibabascbpadreportgettargetreportAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
