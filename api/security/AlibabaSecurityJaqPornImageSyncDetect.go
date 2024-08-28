package security

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// AlibabaSecurityJaqPornImageSyncDetect 聚安全智能鉴黄同步检测接口
// alibaba.security.jaq.porn.image.sync.detect
//
// 同步黄图图像检测接口
func AlibabaSecurityJaqPornImageSyncDetect(ctx context.Context, clt *core.SDKClient, req *security.AlibabaSecurityJaqPornImageSyncDetectAPIRequest, resp *security.AlibabaSecurityJaqPornImageSyncDetectAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
