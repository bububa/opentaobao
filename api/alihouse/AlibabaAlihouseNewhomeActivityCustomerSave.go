package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeActivityCustomerSave 销售活动批量保存定向用户
// alibaba.alihouse.newhome.activity.customer.save
//
// 销售活动批量保存定向用户
func AlibabaAlihouseNewhomeActivityCustomerSave(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeActivityCustomerSaveAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeActivityCustomerSaveAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
