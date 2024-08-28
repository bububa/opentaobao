package vaccin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/vaccin"
)

// AlibabaAlihealthVaccineTradeSubscribeDetailGet 私立疫苗交易-预约详情获取
// alibaba.alihealth.vaccine.trade.subscribe.detail.get
//
// 私立疫苗交易-预约详情获取
func AlibabaAlihealthVaccineTradeSubscribeDetailGet(ctx context.Context, clt *core.SDKClient, req *vaccin.AlibabaAlihealthVaccineTradeSubscribeDetailGetAPIRequest, resp *vaccin.AlibabaAlihealthVaccineTradeSubscribeDetailGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
