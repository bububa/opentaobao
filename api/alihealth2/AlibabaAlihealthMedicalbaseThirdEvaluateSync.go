package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthMedicalbaseThirdEvaluateSync 三方评论信息同步
// alibaba.alihealth.medicalbase.third.evaluate.sync
//
// 三方评论信息同步
func AlibabaAlihealthMedicalbaseThirdEvaluateSync(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthMedicalbaseThirdEvaluateSyncAPIRequest, resp *alihealth2.AlibabaAlihealthMedicalbaseThirdEvaluateSyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
