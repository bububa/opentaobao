package aliospay

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliospay"
)

// Aliyunaliospayrefund 退款接口
// aliyun.alios.pay.refund
//
// 商户用来发起退款的接口
func Aliyunaliospayrefund(clt *core.SDKClient, req *aliospay.AliyunaliospayrefundAPIRequest, session string) (*aliospay.AliyunaliospayrefundAPIResponse, error) {
	var resp aliospay.AliyunaliospayrefundAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
