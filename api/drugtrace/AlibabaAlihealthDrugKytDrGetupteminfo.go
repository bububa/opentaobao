package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytDrGetupteminfo 获取上游温度信息（疫苗）
// alibaba.alihealth.drug.kyt.dr.getupteminfo
//
// 根据追溯码及企业ID获取上游运输及存储温度信息（疫苗）
func AlibabaAlihealthDrugKytDrGetupteminfo(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytDrGetupteminfoAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytDrGetupteminfoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
