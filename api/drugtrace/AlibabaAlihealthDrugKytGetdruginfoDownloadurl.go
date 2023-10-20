package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytGetdruginfoDownloadurl 药品全量数据下载
// alibaba.alihealth.drug.kyt.getdruginfo.downloadurl
//
// 药品全量数据下载
func AlibabaAlihealthDrugKytGetdruginfoDownloadurl(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytGetdruginfoDownloadurlAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytGetdruginfoDownloadurlAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
