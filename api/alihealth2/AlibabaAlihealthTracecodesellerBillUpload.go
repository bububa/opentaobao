package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// Alibabaalihealthtracecodesellerbillupload 上传入出库单api
// alibaba.alihealth.tracecodeseller.bill.upload
//
// 上传入出库单api
func Alibabaalihealthtracecodesellerbillupload(clt *core.SDKClient, req *alihealth2.AlibabaalihealthtracecodesellerbilluploadAPIRequest, session string) (*alihealth2.AlibabaalihealthtracecodesellerbilluploadAPIResponse, error) {
	var resp alihealth2.AlibabaalihealthtracecodesellerbilluploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
