package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

/* AlibabaAlihealthDrugKytFiledownload
处理失败单据下载
alibaba.alihealth.drug.kyt.filedownload

处理失败单据下载 */
func AlibabaAlihealthDrugKytFiledownload(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytFiledownloadAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytFiledownloadAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugKytFiledownloadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
