package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// Alibabamoshrbackgroundreportnotify 背调公司背调结果通知
// alibaba.mos.hr.background.report.notify
//
// 背调公司背调结果通知
func Alibabamoshrbackgroundreportnotify(clt *core.SDKClient, req *mos.AlibabamoshrbackgroundreportnotifyAPIRequest, session string) (*mos.AlibabamoshrbackgroundreportnotifyAPIResponse, error) {
	var resp mos.AlibabamoshrbackgroundreportnotifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
