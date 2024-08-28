package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugCodeKytYyApplycode 医院药品子码申请接口
// alibaba.alihealth.drug.code.kyt.yy.applycode
//
// 根据父码及所属企业ID生成子码信息
func AlibabaAlihealthDrugCodeKytYyApplycode(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugCodeKytYyApplycodeAPIRequest, resp *drugtrace.AlibabaAlihealthDrugCodeKytYyApplycodeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
