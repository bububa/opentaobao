package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugCodeKytSpeciaVaccinQuerycode 根据码查询码信息（疫苗）
// alibaba.alihealth.drug.code.kyt.specia.vaccin.querycode
//
// 服务描述
// 此接口，针对有码药品，提供可通过追溯码获取该药品的基础信息和生产状况信息。
// 核查平台优先过滤非8开头的，长度非20位数字的码信息。
func AlibabaAlihealthDrugCodeKytSpeciaVaccinQuerycode(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugCodeKytSpeciaVaccinQuerycodeAPIRequest, resp *drugtrace.AlibabaAlihealthDrugCodeKytSpeciaVaccinQuerycodeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
