package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugDownloadGetentdailytaskdtolist 码上放心数据落地-获取每天日报
// alibaba.alihealth.drug.download.getentdailytaskdtolist
//
// 码上放心数据落地-获取每天日报
func AlibabaAlihealthDrugDownloadGetentdailytaskdtolist(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugDownloadGetentdailytaskdtolistAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugDownloadGetentdailytaskdtolistAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugDownloadGetentdailytaskdtolistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
