package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytDrAssociateequi 疫苗单据与设备绑定
// alibaba.alihealth.drug.kyt.dr.associateequi
//
// 疫苗单据与设备绑定
func AlibabaAlihealthDrugKytDrAssociateequi(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytDrAssociateequiAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytDrAssociateequiAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
