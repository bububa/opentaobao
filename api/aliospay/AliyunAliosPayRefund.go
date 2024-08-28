package aliospay

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliospay"
)

// AliyunAliosPayRefund 退款接口
// aliyun.alios.pay.refund
//
// 商户用来发起退款的接口
func AliyunAliosPayRefund(ctx context.Context, clt *core.SDKClient, req *aliospay.AliyunAliosPayRefundAPIRequest, resp *aliospay.AliyunAliosPayRefundAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
