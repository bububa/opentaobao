package security

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// AlibabaSecurityJaqOcrImageSyncDetect 聚安全图文识别同步检测接口
// alibaba.security.jaq.ocr.image.sync.detect
//
// 图像字符识别同步检测接口
func AlibabaSecurityJaqOcrImageSyncDetect(clt *core.SDKClient, req *security.AlibabaSecurityJaqOcrImageSyncDetectAPIRequest, session string) (*security.AlibabaSecurityJaqOcrImageSyncDetectAPIResponse, error) {
	var resp security.AlibabaSecurityJaqOcrImageSyncDetectAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
