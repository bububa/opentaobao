package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytDrTransportupload 疫苗运输温度上传
// alibaba.alihealth.drug.kyt.dr.transportupload
//
// 疫苗运输温度上传
func AlibabaAlihealthDrugKytDrTransportupload(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytDrTransportuploadAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytDrTransportuploadAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugKytDrTransportuploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
