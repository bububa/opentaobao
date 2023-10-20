package shenjing

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/shenjing"
)

// AlibabaIbShenjingVisitorPadUploadface 访客PAD上传人脸
// alibaba.ib.shenjing.visitor.pad.uploadface
//
// 访客PAD端上传人脸。
func AlibabaIbShenjingVisitorPadUploadface(clt *core.SDKClient, req *shenjing.AlibabaIbShenjingVisitorPadUploadfaceAPIRequest, resp *shenjing.AlibabaIbShenjingVisitorPadUploadfaceAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
