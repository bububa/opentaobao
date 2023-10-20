package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytremnantbillupload 零头出入库单据上传
// alibaba.alihealth.drug.kyt.remnantbill.upload
//
// 零头出入库单据上传
func Alibabaalihealthdrugkytremnantbillupload(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytremnantbilluploadAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytremnantbilluploadAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytremnantbilluploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
