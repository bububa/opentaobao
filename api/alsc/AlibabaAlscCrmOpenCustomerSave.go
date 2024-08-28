package alsc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmOpenCustomerSave 保存和更新顾客
// alibaba.alsc.crm.open.customer.save
//
// 用来保存顾客，如果已经存在的话，则更新顾客
func AlibabaAlscCrmOpenCustomerSave(ctx context.Context, clt *core.SDKClient, req *alsc.AlibabaAlscCrmOpenCustomerSaveAPIRequest, resp *alsc.AlibabaAlscCrmOpenCustomerSaveAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
