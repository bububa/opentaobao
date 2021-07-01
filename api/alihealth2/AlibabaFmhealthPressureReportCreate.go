package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

/* AlibabaFmhealthPressureReportCreate
血压报告接口
alibaba.fmhealth.pressure.report.create

生成用户血压测量报告 */
func AlibabaFmhealthPressureReportCreate(clt *core.SDKClient, req *alihealth2.AlibabaFmhealthPressureReportCreateAPIRequest, session string) (*alihealth2.AlibabaFmhealthPressureReportCreateAPIResponse, error) {
	var resp alihealth2.AlibabaFmhealthPressureReportCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
