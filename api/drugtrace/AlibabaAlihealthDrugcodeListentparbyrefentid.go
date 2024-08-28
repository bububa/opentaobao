package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugcodeListentparbyrefentid 根据企业id获取往来单位
// alibaba.alihealth.drugcode.listentparbyrefentid
//
// 根据企业id获取往来单位
func AlibabaAlihealthDrugcodeListentparbyrefentid(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugcodeListentparbyrefentidAPIRequest, resp *drugtrace.AlibabaAlihealthDrugcodeListentparbyrefentidAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
