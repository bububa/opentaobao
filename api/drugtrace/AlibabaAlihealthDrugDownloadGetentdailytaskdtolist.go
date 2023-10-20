package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugDownloadGetentdailytaskdtolist 码上放心数据落地-获取每天日报
// alibaba.alihealth.drug.download.getentdailytaskdtolist
//
// 码上放心数据落地-获取每天日报
func AlibabaAlihealthDrugDownloadGetentdailytaskdtolist(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugDownloadGetentdailytaskdtolistAPIRequest, resp *drugtrace.AlibabaAlihealthDrugDownloadGetentdailytaskdtolistAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
