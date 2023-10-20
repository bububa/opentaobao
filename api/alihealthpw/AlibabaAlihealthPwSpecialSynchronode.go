package alihealthpw

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthpw"
)

// AlibabaAlihealthPwSpecialSynchronode 合作方同步状态至阿里健康
// alibaba.alihealth.pw.special.synchronode
//
// 合作方同步状态至阿里健康
func AlibabaAlihealthPwSpecialSynchronode(clt *core.SDKClient, req *alihealthpw.AlibabaAlihealthPwSpecialSynchronodeAPIRequest, resp *alihealthpw.AlibabaAlihealthPwSpecialSynchronodeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
