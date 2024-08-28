package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytGetentlicense 获取企业资质
// alibaba.alihealth.drug.kyt.getentlicense
//
// 获取企业的资质信息。
func AlibabaAlihealthDrugKytGetentlicense(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytGetentlicenseAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytGetentlicenseAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
