package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugCodeKytHospitalsenddrugmachine 医院发药机
// alibaba.alihealth.drug.code.kyt.hospitalsenddrugmachine
//
// 此接口针对有码药品，提供可通过追溯码获取该药品的基础信息和生产信息；
// 核查平台优先过滤非8开头的，长度非20位数字的码信息。
// 提供给医院发药机使用
func AlibabaAlihealthDrugCodeKytHospitalsenddrugmachine(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugCodeKytHospitalsenddrugmachineAPIRequest, resp *drugtrace.AlibabaAlihealthDrugCodeKytHospitalsenddrugmachineAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
