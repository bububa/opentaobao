package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

/* AlibabaAlihealthTracecodesellerBillUpload
上传入出库单api
alibaba.alihealth.tracecodeseller.bill.upload

上传入出库单api */
func AlibabaAlihealthTracecodesellerBillUpload(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthTracecodesellerBillUploadAPIRequest, session string) (*alihealth2.AlibabaAlihealthTracecodesellerBillUploadAPIResponse, error) {
	var resp alihealth2.AlibabaAlihealthTracecodesellerBillUploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
