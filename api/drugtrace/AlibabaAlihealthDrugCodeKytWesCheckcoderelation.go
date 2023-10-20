package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugcodekytwescheckcoderelation 检查输入的码之间是否有上下级关系
// alibaba.alihealth.drug.code.kyt.wes.checkcoderelation
//
// 检查输入的码之间是否有上下级关系
func Alibabaalihealthdrugcodekytwescheckcoderelation(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugcodekytwescheckcoderelationAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugcodekytwescheckcoderelationAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugcodekytwescheckcoderelationAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
