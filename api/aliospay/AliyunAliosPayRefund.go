package aliospay

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliospay"
)

// AliyunAliosPayRefund 退款接口
// aliyun.alios.pay.refund
//
// 商户用来发起退款的接口
func AliyunAliosPayRefund(clt *core.SDKClient, req *aliospay.AliyunAliosPayRefundAPIRequest, session string) (*aliospay.AliyunAliosPayRefundAPIResponse, error) {
	var resp aliospay.AliyunAliosPayRefundAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
