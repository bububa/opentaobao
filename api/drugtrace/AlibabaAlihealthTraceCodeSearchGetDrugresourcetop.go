package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthTraceCodeSearchGetDrugresourcetop 根据码获取码信息
// alibaba.alihealth.trace.code.search.get.drugresourcetop
//
// 根据码获取码信息
func AlibabaAlihealthTraceCodeSearchGetDrugresourcetop(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthTraceCodeSearchGetDrugresourcetopAPIRequest, resp *drugtrace.AlibabaAlihealthTraceCodeSearchGetDrugresourcetopAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
