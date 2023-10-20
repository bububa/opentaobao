package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeActivityDepositUnbind 销售活动解绑预存金商品
// alibaba.alihouse.newhome.activity.deposit.unbind
//
// 销售活动解绑预存金商品
func AlibabaAlihouseNewhomeActivityDepositUnbind(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeActivityDepositUnbindAPIRequest, session string) (*alihouse.AlibabaAlihouseNewhomeActivityDepositUnbindAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseNewhomeActivityDepositUnbindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
