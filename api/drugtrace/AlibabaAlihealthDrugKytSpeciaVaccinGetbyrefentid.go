package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytSpeciaVaccinGetbyrefentid 根据企业唯一标识查看企业详情
// alibaba.alihealth.drug.kyt.specia.vaccin.getbyrefentid
//
// 根据企业唯一标识查看企业详细信息
func AlibabaAlihealthDrugKytSpeciaVaccinGetbyrefentid(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytSpeciaVaccinGetbyrefentidAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytSpeciaVaccinGetbyrefentidAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
