package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeActivityCustomerSave 销售活动批量保存定向用户
// alibaba.alihouse.newhome.activity.customer.save
//
// 销售活动批量保存定向用户
func AlibabaAlihouseNewhomeActivityCustomerSave(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeActivityCustomerSaveAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeActivityCustomerSaveAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
