package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytDrStorageupload 疫苗存储温度上传
// alibaba.alihealth.drug.kyt.dr.storageupload
//
// 疫苗存储温度上传
func AlibabaAlihealthDrugKytDrStorageupload(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytDrStorageuploadAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytDrStorageuploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
