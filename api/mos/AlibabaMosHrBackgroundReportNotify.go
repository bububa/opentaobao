package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// AlibabaMosHrBackgroundReportNotify 背调公司背调结果通知
// alibaba.mos.hr.background.report.notify
//
// 背调公司背调结果通知
func AlibabaMosHrBackgroundReportNotify(clt *core.SDKClient, req *mos.AlibabaMosHrBackgroundReportNotifyAPIRequest, session string) (*mos.AlibabaMosHrBackgroundReportNotifyAPIResponse, error) {
	var resp mos.AlibabaMosHrBackgroundReportNotifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
