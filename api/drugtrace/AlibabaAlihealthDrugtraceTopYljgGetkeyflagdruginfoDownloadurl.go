package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurl 获取重点追溯品种明细下载URL
// alibaba.alihealth.drugtrace.top.yljg.getkeyflagdruginfo.downloadurl
//
// 获取重点追溯品种明细下载URL
func AlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurl(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlAPIRequest, resp *drugtrace.AlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
