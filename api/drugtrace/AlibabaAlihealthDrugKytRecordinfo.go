package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytRecordinfo 快易通健康检查
// alibaba.alihealth.drug.kyt.recordinfo
//
// 快易通健康检查
func AlibabaAlihealthDrugKytRecordinfo(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytRecordinfoAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytRecordinfoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
