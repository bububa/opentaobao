package alihealth

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth"
)

// AlibabaAlihealthPrescriptionAuthGet 阿里健康处方平台获取授权码
// alibaba.alihealth.prescription.auth.get
//
// 获取处方用户授权
func AlibabaAlihealthPrescriptionAuthGet(clt *core.SDKClient, req *alihealth.AlibabaAlihealthPrescriptionAuthGetAPIRequest, resp *alihealth.AlibabaAlihealthPrescriptionAuthGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
