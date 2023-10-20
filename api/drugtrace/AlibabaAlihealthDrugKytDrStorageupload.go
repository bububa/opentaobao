package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytdrstorageupload 疫苗存储温度上传
// alibaba.alihealth.drug.kyt.dr.storageupload
//
// 疫苗存储温度上传
func Alibabaalihealthdrugkytdrstorageupload(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytdrstorageuploadAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytdrstorageuploadAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytdrstorageuploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
