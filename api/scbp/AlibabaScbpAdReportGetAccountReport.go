package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpadreportgetaccountreport 账户报告
// alibaba.scbp.ad.report.get.account.report
//
// 账户报告
func Alibabascbpadreportgetaccountreport(clt *core.SDKClient, req *scbp.AlibabascbpadreportgetaccountreportAPIRequest, session string) (*scbp.AlibabascbpadreportgetaccountreportAPIResponse, error) {
	var resp scbp.AlibabascbpadreportgetaccountreportAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
