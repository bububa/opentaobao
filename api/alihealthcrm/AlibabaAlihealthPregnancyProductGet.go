package alihealthcrm

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthcrm"
)

// AlibabaAlihealthPregnancyProductGet 备孕首页获取达人配置商品
// alibaba.alihealth.pregnancy.product.get
//
// 备孕首页获取达人配置商品
func AlibabaAlihealthPregnancyProductGet(ctx context.Context, clt *core.SDKClient, req *alihealthcrm.AlibabaAlihealthPregnancyProductGetAPIRequest, resp *alihealthcrm.AlibabaAlihealthPregnancyProductGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
