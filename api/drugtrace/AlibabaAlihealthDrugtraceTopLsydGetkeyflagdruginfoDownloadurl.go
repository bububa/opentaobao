package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurl 获取重点追溯品种明细下载URL
// alibaba.alihealth.drugtrace.top.lsyd.getkeyflagdruginfo.downloadurl
//
// 获取重点追溯品种明细下载URL
func AlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurl(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlAPIRequest, resp *drugtrace.AlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
