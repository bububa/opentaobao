package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugcodecodecheckhospital 码核查状态同步-医院
// alibaba.alihealth.drug.code.code.check.hospital
//
// 码核查状态同步-医院
func Alibabaalihealthdrugcodecodecheckhospital(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugcodecodecheckhospitalAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugcodecodecheckhospitalAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugcodecodecheckhospitalAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
