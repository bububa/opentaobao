package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

/* AlibabaAlihealthDrugDownloadGetentauthent
获取授权企业列表
alibaba.alihealth.drug.download.getentauthent

D2D数据落地获取授权企业列表 */
func AlibabaAlihealthDrugDownloadGetentauthent(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugDownloadGetentauthentAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugDownloadGetentauthentAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugDownloadGetentauthentAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
