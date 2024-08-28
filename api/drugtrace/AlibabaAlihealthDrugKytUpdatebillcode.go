package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytUpdatebillcode 零售修改出入库单追溯码
// alibaba.alihealth.drug.kyt.updatebillcode
//
// 零售修改出入库单追溯码
func AlibabaAlihealthDrugKytUpdatebillcode(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytUpdatebillcodeAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytUpdatebillcodeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
