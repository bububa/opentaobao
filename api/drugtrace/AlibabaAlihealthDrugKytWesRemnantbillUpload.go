package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytwesremnantbillupload wes零头出入库单据上传
// alibaba.alihealth.drug.kyt.wes.remnantbill.upload
//
// wes零头出入库单据上传
func Alibabaalihealthdrugkytwesremnantbillupload(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytwesremnantbilluploadAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytwesremnantbilluploadAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytwesremnantbilluploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
