package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugCodeKytWesCheckcoderelation 检查输入的码之间是否有上下级关系
// alibaba.alihealth.drug.code.kyt.wes.checkcoderelation
//
// 检查输入的码之间是否有上下级关系
func AlibabaAlihealthDrugCodeKytWesCheckcoderelation(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugCodeKytWesCheckcoderelationAPIRequest, resp *drugtrace.AlibabaAlihealthDrugCodeKytWesCheckcoderelationAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
