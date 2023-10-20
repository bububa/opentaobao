package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugDownloadEntlist 企业下载列表
// alibaba.alihealth.drug.download.entlist
//
// 获取企业的下载文件列表
func AlibabaAlihealthDrugDownloadEntlist(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugDownloadEntlistAPIRequest, resp *drugtrace.AlibabaAlihealthDrugDownloadEntlistAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
