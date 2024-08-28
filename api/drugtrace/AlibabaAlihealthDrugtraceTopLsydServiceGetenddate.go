package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugtraceTopLsydServiceGetenddate 获取企业服务截止时间
// alibaba.alihealth.drugtrace.top.lsyd.service.getenddate
//
// 获取企业服务截止时间
func AlibabaAlihealthDrugtraceTopLsydServiceGetenddate(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugtraceTopLsydServiceGetenddateAPIRequest, resp *drugtrace.AlibabaAlihealthDrugtraceTopLsydServiceGetenddateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
