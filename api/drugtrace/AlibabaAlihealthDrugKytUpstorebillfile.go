package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytUpstorebillfile 上传零售出入库单(上传文件)
// alibaba.alihealth.drug.kyt.upstorebillfile
//
// 上传零售出入库单(上传文件)
func AlibabaAlihealthDrugKytUpstorebillfile(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytUpstorebillfileAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytUpstorebillfileAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
