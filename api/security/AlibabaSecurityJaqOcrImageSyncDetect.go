package security

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// AlibabaSecurityJaqOcrImageSyncDetect 聚安全图文识别同步检测接口
// alibaba.security.jaq.ocr.image.sync.detect
//
// 图像字符识别同步检测接口
func AlibabaSecurityJaqOcrImageSyncDetect(clt *core.SDKClient, req *security.AlibabaSecurityJaqOcrImageSyncDetectAPIRequest, resp *security.AlibabaSecurityJaqOcrImageSyncDetectAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
