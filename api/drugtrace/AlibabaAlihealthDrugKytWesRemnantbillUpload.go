package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytWesRemnantbillUpload wes零头出入库单据上传
// alibaba.alihealth.drug.kyt.wes.remnantbill.upload
//
// wes零头出入库单据上传
func AlibabaAlihealthDrugKytWesRemnantbillUpload(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytWesRemnantbillUploadAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytWesRemnantbillUploadAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
