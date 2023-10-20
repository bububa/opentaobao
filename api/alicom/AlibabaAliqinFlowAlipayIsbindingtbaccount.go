package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// Alibabaaliqinflowalipayisbindingtbaccount 判断支付宝用户是否绑定淘宝账号
// alibaba.aliqin.flow.alipay.isbindingtbaccount
//
// 判断支付宝用户是否绑定淘宝账号
func Alibabaaliqinflowalipayisbindingtbaccount(clt *core.SDKClient, req *alicom.AlibabaaliqinflowalipayisbindingtbaccountAPIRequest, session string) (*alicom.AlibabaaliqinflowalipayisbindingtbaccountAPIResponse, error) {
	var resp alicom.AlibabaaliqinflowalipayisbindingtbaccountAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
