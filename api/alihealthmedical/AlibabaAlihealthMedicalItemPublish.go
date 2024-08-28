package alihealthmedical

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthmedical"
)

// AlibabaAlihealthMedicalItemPublish 三方入驻-开通服务
// alibaba.alihealth.medical.item.publish
//
// 三方入驻-开通服务
func AlibabaAlihealthMedicalItemPublish(ctx context.Context, clt *core.SDKClient, req *alihealthmedical.AlibabaAlihealthMedicalItemPublishAPIRequest, resp *alihealthmedical.AlibabaAlihealthMedicalItemPublishAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
