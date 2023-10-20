package aliospay

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliospay"
)

// AliyunAliosPayTokenGet 获取支付token
// aliyun.alios.pay.token.get
//
// 商户用来获取支付的授权token
func AliyunAliosPayTokenGet(clt *core.SDKClient, req *aliospay.AliyunAliosPayTokenGetAPIRequest, resp *aliospay.AliyunAliosPayTokenGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
