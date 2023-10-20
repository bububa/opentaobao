package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomerightbindback 权限回流
// alibaba.alihouse.newhome.right.bind.back
//
// 权限回流
func Alibabaalihousenewhomerightbindback(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomerightbindbackAPIRequest, session string) (*alihouse.AlibabaalihousenewhomerightbindbackAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomerightbindbackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
