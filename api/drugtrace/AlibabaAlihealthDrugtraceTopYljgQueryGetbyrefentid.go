package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentid 根据企业唯一标识查看企业详细信息
// alibaba.alihealth.drugtrace.top.yljg.query.getbyrefentid
//
// 根据企业唯一标识查看企业详细信息
func AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentid(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidAPIRequest, resp *drugtrace.AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
