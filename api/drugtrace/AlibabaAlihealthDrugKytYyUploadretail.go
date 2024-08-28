package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytYyUploadretail 医院上传出库信息
// alibaba.alihealth.drug.kyt.yy.uploadretail
//
// 医院上传出库信息接口
func AlibabaAlihealthDrugKytYyUploadretail(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytYyUploadretailAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytYyUploadretailAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
