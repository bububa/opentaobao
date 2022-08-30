package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytGetdruginfoDownloadurl 药品全量数据下载
// alibaba.alihealth.drug.kyt.getdruginfo.downloadurl
//
// 药品全量数据下载
func AlibabaAlihealthDrugKytGetdruginfoDownloadurl(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytGetdruginfoDownloadurlAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytGetdruginfoDownloadurlAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugKytGetdruginfoDownloadurlAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
