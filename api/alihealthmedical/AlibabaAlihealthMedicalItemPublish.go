package alihealthmedical

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthmedical"
)

// AlibabaAlihealthMedicalItemPublish 三方入驻-开通服务
// alibaba.alihealth.medical.item.publish
//
// 三方入驻-开通服务
func AlibabaAlihealthMedicalItemPublish(clt *core.SDKClient, req *alihealthmedical.AlibabaAlihealthMedicalItemPublishAPIRequest, resp *alihealthmedical.AlibabaAlihealthMedicalItemPublishAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
