package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugCodeKytWesCheckcoderelation 检查输入的码之间是否有上下级关系
// alibaba.alihealth.drug.code.kyt.wes.checkcoderelation
//
// 检查输入的码之间是否有上下级关系
func AlibabaAlihealthDrugCodeKytWesCheckcoderelation(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugCodeKytWesCheckcoderelationAPIRequest, resp *drugtrace.AlibabaAlihealthDrugCodeKytWesCheckcoderelationAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
