package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurl 获取重点追溯品种明细下载URL
// alibaba.alihealth.drugtrace.top.yljg.getkeyflagdruginfo.downloadurl
//
// 获取重点追溯品种明细下载URL
func AlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurl(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlAPIRequest, resp *drugtrace.AlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
