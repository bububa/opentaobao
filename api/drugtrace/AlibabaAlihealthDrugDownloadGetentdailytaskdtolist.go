package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugDownloadGetentdailytaskdtolist 码上放心数据落地-获取每天日报
// alibaba.alihealth.drug.download.getentdailytaskdtolist
//
// 码上放心数据落地-获取每天日报
func AlibabaAlihealthDrugDownloadGetentdailytaskdtolist(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugDownloadGetentdailytaskdtolistAPIRequest, resp *drugtrace.AlibabaAlihealthDrugDownloadGetentdailytaskdtolistAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
