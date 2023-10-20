package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugcodeadvancebillflowdirection 单据流向查询
// alibaba.alihealth.drug.code.advance.bill.flow.direction
//
// 单据流向查询
func Alibabaalihealthdrugcodeadvancebillflowdirection(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugcodeadvancebillflowdirectionAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugcodeadvancebillflowdirectionAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugcodeadvancebillflowdirectionAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
