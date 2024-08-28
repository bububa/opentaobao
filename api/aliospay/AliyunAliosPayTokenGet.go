package aliospay

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliospay"
)

// AliyunAliosPayTokenGet 获取支付token
// aliyun.alios.pay.token.get
//
// 商户用来获取支付的授权token
func AliyunAliosPayTokenGet(ctx context.Context, clt *core.SDKClient, req *aliospay.AliyunAliosPayTokenGetAPIRequest, resp *aliospay.AliyunAliosPayTokenGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
