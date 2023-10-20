package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthDocbaseUserinfoAlipayidGet 根据健康ID获取支付宝ID
// alibaba.alihealth.docbase.userinfo.alipayid.get
//
// 根据健康ID获取支付宝ID
func AlibabaAlihealthDocbaseUserinfoAlipayidGet(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthDocbaseUserinfoAlipayidGetAPIRequest, resp *alihealth2.AlibabaAlihealthDocbaseUserinfoAlipayidGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
