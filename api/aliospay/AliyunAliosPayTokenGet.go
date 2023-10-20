package aliospay

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliospay"
)

// Aliyunaliospaytokenget 获取支付token
// aliyun.alios.pay.token.get
//
// 商户用来获取支付的授权token
func Aliyunaliospaytokenget(clt *core.SDKClient, req *aliospay.AliyunaliospaytokengetAPIRequest, session string) (*aliospay.AliyunaliospaytokengetAPIResponse, error) {
	var resp aliospay.AliyunaliospaytokengetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
