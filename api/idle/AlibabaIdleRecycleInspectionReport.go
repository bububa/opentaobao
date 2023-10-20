package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// Alibabaidlerecycleinspectionreport 鉴定报告
// alibaba.idle.recycle.inspection.report
//
// 回收商鉴定报告
func Alibabaidlerecycleinspectionreport(clt *core.SDKClient, req *idle.AlibabaidlerecycleinspectionreportAPIRequest, session string) (*idle.AlibabaidlerecycleinspectionreportAPIResponse, error) {
	var resp idle.AlibabaidlerecycleinspectionreportAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
