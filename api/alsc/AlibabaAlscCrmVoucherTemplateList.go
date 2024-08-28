package alsc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmVoucherTemplateList 获取优惠券模版列表
// alibaba.alsc.crm.voucher.template.list
//
// 获取优惠券模版列表
func AlibabaAlscCrmVoucherTemplateList(ctx context.Context, clt *core.SDKClient, req *alsc.AlibabaAlscCrmVoucherTemplateListAPIRequest, resp *alsc.AlibabaAlscCrmVoucherTemplateListAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
