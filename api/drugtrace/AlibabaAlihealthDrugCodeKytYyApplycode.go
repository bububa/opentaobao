package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugCodeKytYyApplycode 医院药品子码申请接口
// alibaba.alihealth.drug.code.kyt.yy.applycode
//
// 根据父码及所属企业ID生成子码信息
func AlibabaAlihealthDrugCodeKytYyApplycode(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugCodeKytYyApplycodeAPIRequest, resp *drugtrace.AlibabaAlihealthDrugCodeKytYyApplycodeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
