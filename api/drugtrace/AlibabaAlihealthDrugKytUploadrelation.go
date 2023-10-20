package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytUploadrelation 关联关系上传
// alibaba.alihealth.drug.kyt.uploadrelation
//
// 关联关系上传
func AlibabaAlihealthDrugKytUploadrelation(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytUploadrelationAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytUploadrelationAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
