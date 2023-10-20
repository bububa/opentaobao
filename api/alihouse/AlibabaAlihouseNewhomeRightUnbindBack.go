package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomerightunbindback 权限回流-解绑
// alibaba.alihouse.newhome.right.unbind.back
//
// 权限回流-解绑
func Alibabaalihousenewhomerightunbindback(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomerightunbindbackAPIRequest, session string) (*alihouse.AlibabaalihousenewhomerightunbindbackAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomerightunbindbackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
