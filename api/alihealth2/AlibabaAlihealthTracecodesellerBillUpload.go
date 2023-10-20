package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthTracecodesellerBillUpload 上传入出库单api
// alibaba.alihealth.tracecodeseller.bill.upload
//
// 上传入出库单api
func AlibabaAlihealthTracecodesellerBillUpload(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthTracecodesellerBillUploadAPIRequest, resp *alihealth2.AlibabaAlihealthTracecodesellerBillUploadAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
