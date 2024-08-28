package security

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// AlibabaSecurityJaqRpCloudOcrCheck ocr同时实名校验
// alibaba.security.jaq.rp.cloud.ocr.check
//
// 聚安全实人认证证件OCR识别功能接口
func AlibabaSecurityJaqRpCloudOcrCheck(ctx context.Context, clt *core.SDKClient, req *security.AlibabaSecurityJaqRpCloudOcrCheckAPIRequest, resp *security.AlibabaSecurityJaqRpCloudOcrCheckAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
