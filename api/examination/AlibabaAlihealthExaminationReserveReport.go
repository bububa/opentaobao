package examination

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/examination"
)

/* AlibabaAlihealthExaminationReserveReport
体检机构对接_体检报告查询
alibaba.alihealth.examination.reserve.report

体检机构对接_体检报告获取 */
func AlibabaAlihealthExaminationReserveReport(clt *core.SDKClient, req *examination.AlibabaAlihealthExaminationReserveReportAPIRequest, session string) (*examination.AlibabaAlihealthExaminationReserveReportAPIResponse, error) {
	var resp examination.AlibabaAlihealthExaminationReserveReportAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
