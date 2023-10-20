package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytUploadb2cbill 快易通零售B2C
// alibaba.alihealth.drug.kyt.uploadb2cbill
//
// 快易通零售B2C单据上传
func AlibabaAlihealthDrugKytUploadb2cbill(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytUploadb2cbillAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytUploadb2cbillAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugKytUploadb2cbillAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
