package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugCodeCodeCheckHospital 码核查状态同步-医院
// alibaba.alihealth.drug.code.code.check.hospital
//
// 码核查状态同步-医院
func AlibabaAlihealthDrugCodeCodeCheckHospital(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugCodeCodeCheckHospitalAPIRequest, resp *drugtrace.AlibabaAlihealthDrugCodeCodeCheckHospitalAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
