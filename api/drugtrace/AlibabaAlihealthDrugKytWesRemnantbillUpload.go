package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytWesRemnantbillUpload wes零头出入库单据上传
// alibaba.alihealth.drug.kyt.wes.remnantbill.upload
//
// wes零头出入库单据上传
func AlibabaAlihealthDrugKytWesRemnantbillUpload(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytWesRemnantbillUploadAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytWesRemnantbillUploadAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugKytWesRemnantbillUploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
