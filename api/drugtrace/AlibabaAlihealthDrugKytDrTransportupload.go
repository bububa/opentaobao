package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytdrtransportupload 疫苗运输温度上传
// alibaba.alihealth.drug.kyt.dr.transportupload
//
// 疫苗运输温度上传
func Alibabaalihealthdrugkytdrtransportupload(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytdrtransportuploadAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytdrtransportuploadAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytdrtransportuploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
