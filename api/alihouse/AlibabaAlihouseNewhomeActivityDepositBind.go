package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeActivityDepositBind 销售活动绑定预存金商品id
// alibaba.alihouse.newhome.activity.deposit.bind
//
// 销售活动绑定预存金商品id
func AlibabaAlihouseNewhomeActivityDepositBind(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeActivityDepositBindAPIRequest, session string) (*alihouse.AlibabaAlihouseNewhomeActivityDepositBindAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseNewhomeActivityDepositBindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
