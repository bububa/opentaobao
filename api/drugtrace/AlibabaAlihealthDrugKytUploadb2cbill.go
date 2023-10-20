package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytuploadb2cbill 快易通零售B2C
// alibaba.alihealth.drug.kyt.uploadb2cbill
//
// 快易通零售B2C单据上传
func Alibabaalihealthdrugkytuploadb2cbill(clt *core.SDKClient, req *drugtrace.Alibabaalihealthdrugkytuploadb2cbillAPIRequest, session string) (*drugtrace.Alibabaalihealthdrugkytuploadb2cbillAPIResponse, error) {
	var resp drugtrace.Alibabaalihealthdrugkytuploadb2cbillAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
