package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// Alibabaalihealthdocbaseuserinfoalipayidget 根据健康ID获取支付宝ID
// alibaba.alihealth.docbase.userinfo.alipayid.get
//
// 根据健康ID获取支付宝ID
func Alibabaalihealthdocbaseuserinfoalipayidget(clt *core.SDKClient, req *alihealth2.AlibabaalihealthdocbaseuserinfoalipayidgetAPIRequest, session string) (*alihealth2.AlibabaalihealthdocbaseuserinfoalipayidgetAPIResponse, error) {
	var resp alihealth2.AlibabaalihealthdocbaseuserinfoalipayidgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
