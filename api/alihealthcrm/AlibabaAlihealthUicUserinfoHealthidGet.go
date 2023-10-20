package alihealthcrm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthcrm"
)

// AlibabaAlihealthUicUserinfoHealthidGet 获取健康id
// alibaba.alihealth.uic.userinfo.healthid.get
//
// 根据支付宝用户ID获取用户健康ID
func AlibabaAlihealthUicUserinfoHealthidGet(clt *core.SDKClient, req *alihealthcrm.AlibabaAlihealthUicUserinfoHealthidGetAPIRequest, resp *alihealthcrm.AlibabaAlihealthUicUserinfoHealthidGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
