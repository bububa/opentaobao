package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytWesGetdruginfoDownloadurl 药品全量数据下载
// alibaba.alihealth.drug.kyt.wes.getdruginfo.downloadurl
//
// 药品全量数据下载
func AlibabaAlihealthDrugKytWesGetdruginfoDownloadurl(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytWesGetdruginfoDownloadurlAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytWesGetdruginfoDownloadurlAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugKytWesGetdruginfoDownloadurlAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
