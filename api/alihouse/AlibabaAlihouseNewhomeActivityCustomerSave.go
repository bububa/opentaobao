package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeActivityCustomerSave 销售活动批量保存定向用户
// alibaba.alihouse.newhome.activity.customer.save
//
// 销售活动批量保存定向用户
func AlibabaAlihouseNewhomeActivityCustomerSave(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeActivityCustomerSaveAPIRequest, session string) (*alihouse.AlibabaAlihouseNewhomeActivityCustomerSaveAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseNewhomeActivityCustomerSaveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
