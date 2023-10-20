package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdReportGetLastEffectDate 获取最近报表生成时间
// alibaba.scbp.ad.report.get.last.effect.date
//
// 获取最近报表生成时间
func AlibabaScbpAdReportGetLastEffectDate(clt *core.SDKClient, req *scbp.AlibabaScbpAdReportGetLastEffectDateAPIRequest, session string) (*scbp.AlibabaScbpAdReportGetLastEffectDateAPIResponse, error) {
	var resp scbp.AlibabaScbpAdReportGetLastEffectDateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
