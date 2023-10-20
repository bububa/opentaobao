package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytDrStorageupload 疫苗存储温度上传
// alibaba.alihealth.drug.kyt.dr.storageupload
//
// 疫苗存储温度上传
func AlibabaAlihealthDrugKytDrStorageupload(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytDrStorageuploadAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytDrStorageuploadAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugKytDrStorageuploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
