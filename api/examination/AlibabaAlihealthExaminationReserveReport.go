package examination

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/examination"
)

// Alibabaalihealthexaminationreservereport 体检机构对接_体检报告查询
// alibaba.alihealth.examination.reserve.report
//
// 体检机构对接_体检报告获取
func Alibabaalihealthexaminationreservereport(clt *core.SDKClient, req *examination.AlibabaalihealthexaminationreservereportAPIRequest, session string) (*examination.AlibabaalihealthexaminationreservereportAPIResponse, error) {
	var resp examination.AlibabaalihealthexaminationreservereportAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
